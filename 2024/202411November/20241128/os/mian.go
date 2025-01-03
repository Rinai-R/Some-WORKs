package main

import (
	"os"
)

func main() {
	_, err := os.Create("main.go")
	if err != nil {
		panic(err)
	}
	return
}
