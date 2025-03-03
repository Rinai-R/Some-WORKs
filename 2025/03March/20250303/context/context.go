package main

import (
	"context"
	"fmt"
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
	if idx == 5 {
		cancelfunc()
	} else {
		defer cancelfunc()
	}

	MyCTX(newctx)
}

func main() {
	MyCTX(context.Background())
}
