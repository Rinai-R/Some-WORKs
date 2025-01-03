package main

import (
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("d:/test.txt")

	if err == nil {
		fmt.Printf("打开成功，文件内容为：%v", string(content))
	} else {
		fmt.Println("打开失败")
	}
	return
}
