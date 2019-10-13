package tracing

import (
	"fmt"

	"go.opentelemetry.io/exporter/trace/jaeger"
	"go.opentelemetry.io/sdk/export"
)

type JaegerConfig struct {
	Endpoint  string `long:"tracing-jaeger-endpoint"`
	Component string `long:"tracing-jaeger-component" default:"build-tracker"`
}

func (j JaegerConfig) IsConfigured() bool {
	return j.Endpoint != ""
}

func (j JaegerConfig) Exporter() (export.SpanSyncer, error) {
	exporter, err := jaeger.NewExporter(
		jaeger.WithCollectorEndpoint(j.Endpoint),
		jaeger.WithProcess(jaeger.Process{
			ServiceName: "web",
		}),
	)
	if err != nil {
		err = fmt.Errorf("failed to create jaeger exporter: %w", err)
		return nil, err
	}

	return exporter, nil
}
