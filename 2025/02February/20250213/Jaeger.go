package main

import (
	"Golang/2025/02February/20250210/kitex-etcd/App/pkg/opentel"
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func main() {
	sdk, err := opentel.SetupOTelSDK(context.Background(), "Main-Test", "1.0.0")
	if err != nil {
		return
	}
	defer sdk(context.Background())
	ctx := context.Background()

	tracer := otel.Tracer("Main-tracer")
	var span trace.Span
	ctx, span = tracer.Start(ctx, "Main-Test-span")
	defer span.End()
	// 在 span 上设置一些属性
	span.SetAttributes(attribute.String("method", "MyMethod"))

	Next(ctx, 0)
}

func Next(ctx context.Context, depth int) {
	sdk, err := opentel.SetupOTelSDK(context.Background(), "Son", "1.0.0")
	if err != nil {
		return
	}
	defer sdk(context.Background())

	tracer := otel.Tracer("Son-tracer")
	var span trace.Span
	ctx, span = tracer.Start(ctx, "Son-span")
	defer span.End()
	// 在 span 上设置一些属性
	span.SetAttributes(attribute.String("method", "MyMethod"))
	if depth > 10 {
		return
	}
	Next(ctx, depth+1)
}
