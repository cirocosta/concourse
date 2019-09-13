package worker

import (
	"context"
	"os"
	"sync"
	"time"

	"code.cloudfoundry.org/garden"
	"code.cloudfoundry.org/lager"
	"code.cloudfoundry.org/lager/lagerctx"
	"github.com/concourse/concourse/metrics"
)

// containerSweeper is an ifrit.Runner that periodically reports and
// garbage-collects a worker's containers
type containerSweeper struct {
	logger       lager.Logger
	interval     time.Duration
	tsaClient    TSAClient
	gardenClient garden.Client
	maxInFlight  uint16
}

func NewContainerSweeper(
	logger lager.Logger,
	sweepInterval time.Duration,
	tsaClient TSAClient,
	gardenClient garden.Client,
	maxInFlight uint16,
) *containerSweeper {
	return &containerSweeper{
		logger:       logger,
		interval:     sweepInterval,
		tsaClient:    tsaClient,
		gardenClient: gardenClient,
		maxInFlight:  maxInFlight,
	}
}

func (sweeper *containerSweeper) Run(signals <-chan os.Signal, ready chan<- struct{}) error {
	timer := time.NewTicker(sweeper.interval)

	close(ready)

	for {
		select {
		case <-timer.C:
			sweeper.sweep(sweeper.logger.Session("tick"))

		case sig := <-signals:
			sweeper.logger.Info("sweep-cancelled-by-signal", lager.Data{"signal": sig})
			return nil
		}
	}
}

func (sweeper *containerSweeper) sweep(logger lager.Logger) {
	ctx := lagerctx.NewContext(context.Background(), logger)

	containers, err := sweeper.gardenClient.Containers(garden.Properties{})
	if err != nil {
		logger.Error("failed-to-list-containers", err)
	} else {
		metrics.Containers.Set(float64(len(containers)))

		handles := []string{}
		for _, container := range containers {
			handles = append(handles, container.Handle())
		}

		err := sweeper.tsaClient.ReportContainers(ctx, handles)
		if err != nil {
			logger.Error("failed-to-report-containers", err)
		}
	}

	containerHandles, err := sweeper.tsaClient.ContainersToDestroy(ctx)
	if err != nil {
		logger.Error("failed-to-get-containers-to-destroy", err)
	} else {
		var wg sync.WaitGroup
		maxInFlight := make(chan int, sweeper.maxInFlight)

		for _, handle := range containerHandles {
			maxInFlight <- 1
			wg.Add(1)

			go func(handle string) {
				start := time.Now()
				err := sweeper.gardenClient.Destroy(handle)
				if err != nil {
					logger.WithData(lager.Data{"handle": handle}).Error("failed-to-destroy-container", err)
				}

				metrics.
					ContainersSweepingDuration.
					WithLabelValues(metrics.StatusFromError(err)).
					Observe(time.Since(start).Seconds())

				<-maxInFlight
				wg.Done()
			}(handle)
		}
		wg.Wait()
	}
}
