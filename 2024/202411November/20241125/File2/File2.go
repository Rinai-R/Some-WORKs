package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	File, err := os.Open("d:/test.txt")

	if err != nil {
		fmt.Println("打开失败，原因为：", err)
		return
	}

	defer File.Close()

	Reader := bufio.NewReader(File)
	str, err2 := Reader.ReadString('9')
	if err2 == nil {

	}

	fmt.Println(str)

	return
}
