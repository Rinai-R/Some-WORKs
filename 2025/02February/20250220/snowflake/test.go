package main

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"log"
	"sync"
)

func main() {
	n := 10000
	var wg sync.WaitGroup
	wg.Add(n)
	defer wg.Wait()
	node, err := snowflake.NewNode(1) // 1 是节点 ID
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < 10000; i++ {
		go func() {
			fmt.Println(node.Generate())
			defer wg.Done()
		}()
	}
}
