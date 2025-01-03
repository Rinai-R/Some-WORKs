package main

import "fmt"

func ComPare(x string, y string) int {
	charCount := make(map[rune]int)
	for _, char := range x {
		charCount[char]++
	}
	commonCount := 0
	for _, char := range y {
		if charCount[char] > 0 {
			commonCount++
			charCount[char]--
		}
	}

	return commonCount
}

func main() {
	fmt.Println(ComPare("ac", "affsfsaaaaaa"))
}
