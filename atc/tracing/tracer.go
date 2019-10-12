package tracing

import (
	"context"

	"github.com/concourse/concourse/atc/db"
	"go.opentelemetry.io/api/core"
	"go.opentelemetry.io/api/trace"
	"go.opentelemetry.io/sdk/export"
	sdktrace "go.opentelemetry.io/sdk/trace"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 go.opentelemetry.io/api/trace.Tracer
//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 go.opentelemetry.io/api/trace.Span

type Tracer struct {
	Tracer trace.Tracer
}

// GlobalTracer is a tracer that can be accessed anywhere, regardless of whether
// the tracing is enabled or not (if not, no-op will be used under the hood).
//
var GlobalTracer *Tracer

// init initializes the global tracer that can be used by anywhere in the code.
//
// ps.: without `ConfigureTracer`, this is noop.
//
func init() {
	GlobalTracer = &Tracer{
		Tracer: sdktrace.Register(),
	}
}

// BuildRootSpan creates a root span that represents the entire execution of a
// build.
//
// `build` *must not* be nil.
//
func (t *Tracer) BuildRootSpan(build db.Build) trace.Span {
	_, span := t.Tracer.Start(
		context.Background(),
		"build",
	)

	span.SetAttributes(
		core.KeyValue{
			core.Key{"team-name"},
			core.Value{String: build.TeamName()},
		},
		core.KeyValue{
			core.Key{"pipeline-name"},
			core.Value{String: build.PipelineName()},
		},
		core.KeyValue{
			core.Key{"job-name"},
			core.Value{String: build.JobName()},
		},
		core.KeyValue{
			core.Key{"name"},
			core.Value{String: build.Name()},
		},
	)

	return span
}

// ConfigureTracer
//
func ConfigureTracer(exporter export.SpanSyncer) {
	sdktrace.RegisterSpanProcessor(
		sdktrace.NewSimpleSpanProcessor(exporter),
	)
	sdktrace.ApplyConfig(
		sdktrace.Config{
			DefaultSampler: sdktrace.AlwaysSample(),
		},
	)
}
