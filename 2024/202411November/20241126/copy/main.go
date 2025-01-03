package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.ReadFile("d:/test.txt")

	if err != nil {
		fmt.Println("读取失败")
		return
	}

	err2 := os.WriteFile("d:/text2.txt", file, 0666)

	if err2 != nil {
		fmt.Println("写入失败")
		return
	}
	content, err3 := os.ReadFile("d:/text2.txt")

	if err3 == nil {
		fmt.Printf("打开成功，文件内容为：%v", string(content))
	} else {
		fmt.Println("打开失败")
	}

	return

}
