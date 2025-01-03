package main

import (
	"fmt"
	"os"
)

func main() {
	var Name []string

	if len(os.Args) <= 1 {
		fmt.Printf("你没输入名字啊？")
	} else {
		Name = os.Args[1:]
		fmt.Printf("你好！%s", Name)
	}
	return
}
