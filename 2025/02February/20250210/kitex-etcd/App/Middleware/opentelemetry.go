package Middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// OpenTelemetryMiddleware 中间件
func OpenTelemetryMiddleware() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		tracer := otel.Tracer("tracer")
		var span trace.Span
		// 从请求中获取 tracecontext
		c, span = tracer.Start(c, "api-span", trace.WithAttributes(attribute.String("method", string(ctx.Request.Method()))))
		defer span.End()
		// 在 span 上设置一些额外的元数据
		span.SetAttributes(attribute.String("path", string(ctx.Request.Path())))

		// 继续处理请求
		ctx.Next(c)
	}
}
