package main

import (
	"context"
	"fmt"
	"runtime"
)

var idx = 0

func MyCTX(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Context cancelled")
		return
	default:
	}

	fmt.Println("Hello World ", idx)
	idx++

	newctx, cancelfunc := context.WithCancel(ctx)
	defer cancelfunc()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in MyCTX", r)
		}
	}()
	if idx == 5 {
		panic("Context cancelled")
	}

	MyCTX(newctx)
}

func main() {
	MyCTX(context.Background())
	runtime.
}
