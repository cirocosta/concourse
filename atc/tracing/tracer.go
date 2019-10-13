package tracing

import (
	"context"

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

	// Attr TODO
	//
	Attr map[string]string
)

var (
	// GlobalTracer is a tracer that can be accessed anywhere, regardless of whether
	// the tracing is enabled or not (if not, no-op will be used under the hood).
	//
	GlobalTracer *Tracer
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

// LogEvent ... TODO
//
func AddEvent(span trace.Span, msg string, attrs map[string]string) {
	span.AddEvent(context.Background(), msg, keyValueSlice(attrs)...)
}

// StartSpan creates a span, giving back a context that has itself added as the
// parent.
//
// ps.: if `ctx` already has a span set, this span becomes a child of it.
//
func (t *Tracer) StartSpan(
	ctx context.Context,
	component string,
	attrs map[string]string,
) (context.Context, trace.Span) {
	var (
		opts       = []trace.SpanOption{}
		parentSpan = contextSpan(ctx)
	)

	if parentSpan != nil {
		opts = append(opts, trace.ChildOf(parentSpan.SpanContext()))
	}

	_, span := t.Tracer.Start(
		context.Background(),
		component,
		opts...,
	)

	span.SetAttributes(keyValueSlice(attrs)...)
	ctx = withSpan(ctx, span)

	return ctx, span
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

// contextSpan retrieves a parent span from the context.
//
func contextSpan(ctx context.Context) trace.Span {
	span, _ := ctx.Value(contextRootSpanKey{}).(trace.Span)
	return span
}

// withSpan augments a context to have a span set.
//
func withSpan(ctx context.Context, span trace.Span) context.Context {
	if span == nil {
		panic("nil span")
	}

	return context.WithValue(ctx, contextRootSpanKey{}, span)
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
