package tracing

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 go.opentelemetry.io/api/trace.Tracer
//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 go.opentelemetry.io/api/trace.Span

import (
	"context"

	"go.opentelemetry.io/api/trace"
	"go.opentelemetry.io/sdk/export"
	sdktrace "go.opentelemetry.io/sdk/trace"
)

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
	ctx, span := trace.GlobalTracer().Start(
		ctx,
		component,
	)

	span.SetAttributes(keyValueSlice(attrs)...)

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
