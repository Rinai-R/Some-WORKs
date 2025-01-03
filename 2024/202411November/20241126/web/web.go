package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("建立连接中~~~")
	connect, err := net.Dial("tcp", "127.0.0.1:8888")

	if err != nil {
		fmt.Println("链接失败~错误为：", err)
		return
	}
	fmt.Println("链接成功！链接：", connect)

	wg.Add(3)
	reader := bufio.NewReader(os.Stdin)
	go func() {
		for {
			str, err2 := reader.ReadString('\n')
			if err2 != nil {
				fmt.Println("输入失败，错误为：", err2)
			} else {
				n, err3 := connect.Write([]byte(str))

				if err3 != nil {
					fmt.Println("发送失败，错误为：", err3)
					fmt.Println("请重试")
				} else {
					fmt.Printf("你：%v\n一共发送了%d个字节\n", str, n)
				}

			}
			wg.Done()
		}
	}()

	wg.Wait()

}
