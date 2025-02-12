package main

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() (err error) {
	// 优雅地处理SIGINT（CTRL+C）。
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// 设置OpenTelemetry。
	serviceName := "dice"
	serviceVersion := "0.1.0"
	otelShutdown, err := setupOTelSDK(ctx, serviceName, serviceVersion)
	if err != nil {
		return err
	}
	// 适当处理关闭，以避免泄漏。
	defer func() {
		err = errors.Join(err, otelShutdown(context.Background()))
	}()

	// 启动HTTP服务器。
	srv := &http.Server{
		Addr:         ":8080",
		BaseContext:  func(_ net.Listener) context.Context { return ctx },
		ReadTimeout:  time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      newHTTPHandler(),
	}
	srvErr := make(chan error, 1)
	go func() {
		srvErr <- srv.ListenAndServe()
	}()

	// 等待中断。
	select {
	case err = <-srvErr:
		// 启动HTTP服务器时发生错误。
		return err
	case <-ctx.Done():
		// 等待第一次CTRL+C。
		// 尽快停止接收信号通知。
		stop()
	}

	// 当调用Shutdown时，ListenAndServe会立即返回ErrServerClosed。
	err = srv.Shutdown(context.Background())
	return err
}

func newHTTPHandler() http.Handler {
	mux := http.NewServeMux()

	// handleFunc是mux.HandleFunc的替代品
	// 它将处理程序的HTTP仪表与模式作为http.route一起增强。
	handleFunc := func(pattern string, handlerFunc func(http.ResponseWriter, *http.Request)) {
		// 为HTTP仪表配置“http.route”。
		handler := otelhttp.WithRouteTag(pattern, http.HandlerFunc(handlerFunc))
		mux.Handle(pattern, handler)
	}

	// 注册处理程序。
	handleFunc("/rolldice", rolldice)

	// 为整个服务器添加HTTP仪表。
	handler := otelhttp.NewHandler(mux, "/")
	return handler
}
