package k8s_test

import (
	"strings"
	"time"

	"github.com/onsi/gomega/gexec"

	. "github.com/concourse/concourse/topgun"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func WorkerRebalancing(im InstanceManager, fly *Fly) {
	It("eventually has worker connecting to each web nodes over a period of time", func() {
		webInstances := im.GetWebInstances()

		Eventually(func() string {
			workers := fly.GetWorkers()
			Expect(workers).To(HaveLen(1))

			return strings.Split(workers[0].GardenAddress, ":")[0]
		}, 2*time.Minute, 10*time.Second).
			Should(Equal(webInstances[0].GetAddress()))

		Eventually(func() string {
			workers := fly.GetWorkers()

			Expect(workers).To(HaveLen(1))
			return strings.Split(workers[0].GardenAddress, ":")[0]
		}, 2*time.Minute, 10*time.Second).
			Should(Equal(webInstances[1].GetAddress()))
	})
}
