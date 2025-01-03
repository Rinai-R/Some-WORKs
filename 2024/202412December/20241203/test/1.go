package main

import (
	"fmt"
	"time"
)

func main() {
	customTime := time.Date(2024, time.December, 03, 18, 13, 0, 0, time.UTC)
	fmt.Println(customTime)
	fmt.Println(time.Now())
}
