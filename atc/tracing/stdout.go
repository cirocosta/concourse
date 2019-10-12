package tracing

import (
	"fmt"

	"go.opentelemetry.io/exporter/trace/stdout"
	"go.opentelemetry.io/sdk/export"
)

type StdoutConfig struct {
	Configured bool `long:"tracing-stdout"`
}

func (s StdoutConfig) IsConfigured() bool {
	return s.Configured
}

// InitializeStdoutExporter
//
func (s StdoutConfig) Exporter() (export.SpanSyncer, error) {
	exporter, err := stdout.NewExporter(stdout.Options{PrettyPrint: false})
	if err != nil {
		err = fmt.Errorf("failed to create stdout exporter: %w", err)
		return nil, err
	}

	return exporter, nil
}
