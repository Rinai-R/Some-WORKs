package main

import (
	"fmt"
	"net/http"
)

// ping 响应函数
func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong!11111111111111111111111111")
}

func main() {
	http.HandleFunc("/ping", ping)    // 创建路由
	http.ListenAndServe(":8000", nil) // 监听端口及启动服务
}
