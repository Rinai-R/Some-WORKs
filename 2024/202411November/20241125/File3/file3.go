package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 在用户文档路径下创建/打开文件
	file, err := os.OpenFile("C:/Users/chenyue/demo1.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("打开失败:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	if _, err := writer.WriteString("114514+777451 = 2233\n"); err != nil {
		fmt.Println("写入失败，原因为:", err)
		return
	}

	if err := writer.Flush(); err != nil {
		fmt.Println("刷新写入缓冲失败，原因为:", err)
		return
	}

	str, err2 := os.ReadFile("C:/Users/chenyue/demo1.txt")
	if err2 != nil {
		fmt.Println("读取失败，原因为:", err2)
		return
	} else {
		fmt.Println("读取成功！内容：", string(str))
	}
}
