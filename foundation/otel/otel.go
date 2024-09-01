package otel

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.13.0"
)

func NewTracer() (*trace.TracerProvider, error) {
	return initTracer()
}
func initTracer() (*trace.TracerProvider, error) {
	// Create a stdout exporter to export the trace data to the console.

	// exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint("http://localhost:14268/api/traces")))
	// if err != nil {
	// 	log.Fatalf("failed to create Jaeger exporter: %v", err)
	// }

	// tp := tracesdk.NewTracerProvider(
	// 	trace.WithBatcher(exporter),
	// 	trace.WithSampler(trace.TraceIDRatioBased(0.1)),
	// )

	// otel.SetTracerProvider(tp)
	// return tp, nil
	exporter, err := otlptrace.New(
		context.Background(),
		otlptracehttp.NewClient(
			otlptracehttp.WithEndpointURL("http://localhost:14268/api/traces"),
			otlptracehttp.WithInsecure(), // No TLS
		),
	)
	if err != nil {
		return nil, fmt.Errorf("creating new exporter: %w", err)
	}

	// Create a TracerProvider with a batcher and resource attributes
	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithSampler(trace.ParentBased(trace.TraceIDRatioBased(0.5))),
		trace.WithResource(
			resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceNameKey.String("library-app"),
			),
		),
	)

	otel.SetTracerProvider(tp)
	return tp, nil
}

func ShutDownTracer(tp *trace.TracerProvider) { _ = tp.Shutdown(context.Background()) }
