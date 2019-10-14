package tracing

import (
	"fmt"

	"go.opentelemetry.io/exporter/trace/jaeger"
	"go.opentelemetry.io/sdk/export"
)

// Jaeger TODO
//
type Jaeger struct {
	Endpoint  string `long:"tracing-jaeger-endpoint"`
	Component string `long:"tracing-jaeger-component" default:"web"`
}

func (j Jaeger) IsConfigured() bool {
	return j.Endpoint != ""
}

func (j Jaeger) Exporter() (export.SpanSyncer, error) {
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
