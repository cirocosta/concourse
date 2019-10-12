package tracing

import (
	"context"

	"github.com/concourse/concourse/atc/db"
	"go.opentelemetry.io/api/key"
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

var (
	teamAttr     = key.New("team")
	pipelineAttr = key.New("pipeline")
	jobAttr      = key.New("job")
	buildAttr    = key.New("build")
)

// BuildRootSpan creates a root span that represents the entire execution of a
// build.
//
// `build` *must not* be nil.
//
func (t *Tracer) BuildRootSpan(build db.Build) trace.Span {
	const operationName = "build"

	_, span := t.Tracer.Start(
		context.Background(),
		operationName,
	)

	span.SetAttributes(
		teamAttr.String(build.TeamName()),
		pipelineAttr.String(build.PipelineName()),
		jobAttr.String(build.JobName()),
		buildAttr.String(build.Name()),
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
