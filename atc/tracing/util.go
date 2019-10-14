package tracing

import (
	"context"

	"go.opentelemetry.io/api/core"
	"go.opentelemetry.io/api/key"
	"go.opentelemetry.io/api/trace"
)

type (
	// contextRootSpanKey is the key to be used when retrieving or setting values in
	// context objects.
	//
	contextRootSpanKey struct{}
)

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
