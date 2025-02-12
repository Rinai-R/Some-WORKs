package middleware

import (
	"context"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

// OpenTelemetryMiddleware 中间件
func OpenTelemetryMiddleware() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		// 返回一个新的处理函数
		return func(ctx context.Context, req interface{}, resp interface{}) (err error) {
			tracer := otel.Tracer("kitex-tracer")

			// 启动一个新的 span
			ctx, span := tracer.Start(ctx, "kitex-request")
			defer span.End()

			// 在 span 上添加一些信息
			span.SetAttributes(attribute.String("request", "kitex-example"))

			// 调用下一个处理函数
			err = next(ctx, req, resp)
			return err
		}
	}
}
