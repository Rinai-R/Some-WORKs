package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Now()

	fmt.Println(t1.Format("2006年01月02日---15点04分05秒"))
}
