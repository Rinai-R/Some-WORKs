package main

import (
	"fmt"
	"net/http"
)

type MyhelloHandler struct {
	content string
}

func (myhandler *MyhelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, myhandler.content)
}

func main() {
	http.Handle("/", &MyhelloHandler{content: "Ciallo World!"})
	http.ListenAndServe(":8888", nil)

}

// 方式一
//func HelloHandler(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "Ciallo World")
//}
//
//func main() {
//	http.HandleFunc("/", HelloHandler)
//	http.ListenAndServe(":8888", nil)
//}
