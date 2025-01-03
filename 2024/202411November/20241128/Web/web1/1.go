package main

import (
	"fmt"
	"net/http"
)

type myHandler struct {
	content string
}

func (m *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, m.content)
}

func main() {
	http.Handle("/", &myHandler{content: "                      这是一个学生管理系统~\n"})

	http.Handle("/Student/", &myHandler{content: "学号		姓名		性别		绩点\n"})
	http.ListenAndServe(":8080", nil)
}
