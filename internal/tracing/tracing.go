package tracing

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/net/context"
)

func createTracingProvider(url string) (*tracesdk.TracerProvider, error) {
	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	tracingProvider := tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(exporter),
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("my-resto"),
			attribute.String("environment", "staging"),
		)),
	)
	return tracingProvider, nil
}

func Init(url string) error {
	tracingProvider, err := createTracingProvider(url)
	if err != nil {
		return err
	}
	otel.SetTracerProvider(tracingProvider)

	return nil
}

func CreateSpan(ctx context.Context, name string) (context.Context, trace.Span) {
	if ctx == nil {
		ctx = context.Background()
	}
	tracer := otel.Tracer(name)
	ctx, span := tracer.Start(ctx, name)
	return ctx, span
}
