package main

import (
	"fmt"
	"net"
)

func process(connect net.Conn) {
	defer connect.Close()

	for {
		buf := make([]byte, 1024)

		n, err3 := connect.Read(buf)

		if err3 != nil {
			fmt.Println("接受失败，错误原因为：", err3)
			return
		} else {
			fmt.Println("Rinai：", string(buf[0:n]))
		}
	}
}

func main() {
	fmt.Println("连接客户端中....")

	Listen, err := net.Listen("tcp", "127.0.0.1:8888")

	if err != nil {
		fmt.Println("监听失败，错误为：", err)
		return
	}
	for {
		Connect, err1 := Listen.Accept()

		if err1 != nil {
			fmt.Println("连接失败，错误为：", err1)
		} else {
			fmt.Printf("连接成功！Connect=%v,接收到的客户端信息：%v\n", Connect, Connect.RemoteAddr().String())
			go process(Connect)
		}
	}

}
