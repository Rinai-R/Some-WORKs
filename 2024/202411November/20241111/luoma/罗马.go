package main

import (
	"fmt"
)

var Lm = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func LmCount(X string) int {
	var res int
	res = 0
	for i := 0; i < len(X)-1; i++ {
		if Lm[rune(X[i])] >= Lm[rune(X[i+1])] {
			res += Lm[rune(X[i])]
		}
		if Lm[rune(X[i])] < Lm[rune(X[i+1])] {
			res -= Lm[rune(X[i])]
		}
	}
	res += Lm[rune(X[len(X)-1])]
	return res
}

func main() {
	var X string
	fmt.Scanln(&X)
	fmt.Println(LmCount(X))
	return
}
