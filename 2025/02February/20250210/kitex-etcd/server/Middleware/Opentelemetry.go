package Middleware

import (
	"context"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// OpenTelemetryMiddleware 中间件：接受上下文并继续处理
func OpenTelemetryMiddleware() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, req interface{}, resp interface{}) (err error) {
			// 继承 Hertz 传递过来的上下文
			tracer := otel.Tracer("user-tracer")
			var span trace.Span
			ctx, span = tracer.Start(ctx, "user-span")
			defer span.End()
			// 在 span 上设置一些属性
			span.SetAttributes(attribute.String("method", "MyMethod"))

			// 调用下一个处理函数
			err = next(ctx, req, resp)
			return err
		}
	}
}
