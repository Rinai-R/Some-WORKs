package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	FILE, err := os.OpenFile("test.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)

	if err != nil {
		panic(err)
		return
	}
	defer FILE.Close()
	Writer := bufio.NewWriter(FILE)
	t1 := time.Now()
	for i := 1; i < 1000000; i++ {
		_, err = Writer.WriteString("Explosion!")
		if err != nil {
			panic(err)
			return
		}
	}
	err2 := Writer.Flush()
	if err2 != nil {
		panic(err)
		return
	}
	t2 := time.Since(t1)
	fmt.Println(t2)

	t3 := time.Now()

	for i := 1; i < 1000000; i++ {
		_, err3 := FILE.WriteString("Explosion!")
		if err3 != nil {
			panic(err3)
		}
	}
	t4 := time.Since(t3)

	fmt.Println(t4)

}
