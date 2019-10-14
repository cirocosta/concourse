package tracing

import (
	"context"

	"go.opentelemetry.io/api/trace"
	"go.opentelemetry.io/sdk/export"
	sdktrace "go.opentelemetry.io/sdk/trace"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 go.opentelemetry.io/api/trace.Tracer
//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 go.opentelemetry.io/api/trace.Span

type Attrs map[string]string

// StartSpan creates a span, giving back a context that has itself added as the
// parent span.
//
// ps.: if `ctx` already has a span set, this span becomes a child of it.
//
func StartSpan(
	ctx context.Context,
	component string,
	attrs Attrs,
) (context.Context, trace.Span) {
	var (
		opts       = []trace.SpanOption{}
		parentSpan = contextSpan(ctx)
	)

	if parentSpan != nil {
		opts = append(opts, trace.ChildOf(parentSpan.SpanContext()))
	}

	_, span := trace.GlobalTracer().Start(
		context.Background(),
		component,
		opts...,
	)

	span.SetAttributes(keyValueSlice(attrs)...)
	ctx = withSpan(ctx, span)

	return ctx, span
}

// ConfigureTracer configures the sdk to use a given exporter, with a given
// config, making the global tracer not a noop anymore.
//
func ConfigureTracer(exporter export.SpanSyncer) {
	sdktrace.Register()

	sdktrace.RegisterSpanProcessor(
		sdktrace.NewSimpleSpanProcessor(exporter),
	)

	sdktrace.ApplyConfig(
		sdktrace.Config{
			DefaultSampler: sdktrace.AlwaysSample(),
		},
	)
}
