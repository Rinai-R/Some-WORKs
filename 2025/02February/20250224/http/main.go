package main

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := map[interface{}]interface{}{
			"message": "Hello, World!",
			"code":    200,
		}

		// 设置响应头 Content-Type 为 application/json
		w.Header().Set("Content-Type", "application/json")

		// 设置响应状态码
		w.WriteHeader(http.StatusOK)

		// 将结构体序列化为 JSON 并写入响应
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		}
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
