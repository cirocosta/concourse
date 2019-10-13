package tracing

import (
	"context"

	"github.com/concourse/concourse/atc/db"
	"go.opentelemetry.io/api/core"
	"go.opentelemetry.io/api/key"
	"go.opentelemetry.io/api/trace"
	"go.opentelemetry.io/sdk/export"
	sdktrace "go.opentelemetry.io/sdk/trace"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 go.opentelemetry.io/api/trace.Tracer
//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 go.opentelemetry.io/api/trace.Span

type (
	// contextRootSpanKey is the key to be used when retrieving or setting values in
	// context objects.
	//
	contextRootSpanKey struct{}

	// Tracer TODO
	//
	Tracer struct{ Tracer trace.Tracer }
)

var (
	// GlobalTracer is a tracer that can be accessed anywhere, regardless of whether
	// the tracing is enabled or not (if not, no-op will be used under the hood).
	//
	GlobalTracer *Tracer

	// attributes to set in spans
	//
	teamAttr     = key.New("team")
	pipelineAttr = key.New("pipeline")
	jobAttr      = key.New("job")
	buildAttr    = key.New("build")
)

const (
	buildOperationName = "build"
)

// init initializes the global tracer that can be used by anywhere in the code.
//
// ps.: without `ConfigureTracer`, this is noop.
//
func init() {
	GlobalTracer = &Tracer{
		Tracer: sdktrace.Register(),
	}
}

// WithSpan TODO
//
func WithSpan(ctx context.Context, span trace.Span) context.Context {
	if span == nil {
		panic("nil span")
	}

	return context.WithValue(ctx, contextRootSpanKey{}, span)
}

func AddEvent(span trace.Span, msg string, attrs map[string]string) {
	span.AddEvent(context.Background(), msg, keyValueSlice(attrs)...)
}

// contextSpan retrieves a parent span from the context.
//
func contextSpan(ctx context.Context) trace.Span {
	span, _ := ctx.Value(contextRootSpanKey{}).(trace.Span)
	return span
}

func keyValueSlice(attrs map[string]string) []core.KeyValue {
	res := make([]core.KeyValue, len(attrs))

	idx := 0
	for k, v := range attrs {
		res[idx] = key.New(k).String(v)
		idx++
	}

	return res
}

func (t *Tracer) Span(ctx context.Context, stepType string, attrs map[string]string) trace.Span {
	var (
		opts       = []trace.SpanOption{}
		parentSpan = contextSpan(ctx)
	)

	if parentSpan != nil {
		opts = append(opts, trace.ChildOf(parentSpan.SpanContext()))
	}

	_, span := t.Tracer.Start(
		context.Background(),
		stepType,
		opts...,
	)

	span.SetAttributes(keyValueSlice(attrs)...)

	return span
}

// BuildRootSpan creates a root span that represents the entire execution of a
// build.
//
// ps.: `build` *must not* be nil.
//
func (t *Tracer) BuildRootSpan(build db.Build) trace.Span {
	if build == nil {
		panic("nil build")
	}

	_, span := t.Tracer.Start(
		context.Background(),
		buildOperationName,
	)

	span.SetAttributes(
		teamAttr.String(build.TeamName()),
		pipelineAttr.String(build.PipelineName()),
		jobAttr.String(build.JobName()),
		buildAttr.String(build.Name()),
	)

	return span
}

// ConfigureTracer TODO
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
