package main

import (
	"fmt"
	"unicode"
)

const N = 100010

var Op [N]rune
var OpIdx int
var Num [N]float32
var NumIdx int

func OpPush(x rune) {
	Op[OpIdx] = x
	OpIdx++
	return
}
func OpPop() {
	if OpIdx == 0 {
		return
	}
	OpIdx--
	return
}

func NumPush(x float32) {
	Num[NumIdx] = x
	NumIdx++
	return
}
func NumPop() {
	if NumIdx == 0 {
		return
	}
	NumIdx--
	return
}

func eval() {
	if NumIdx >= 2 && OpIdx > 0 {
		var b float32
		b = Num[NumIdx-1]
		NumPop()
		var a float32
		a = Num[NumIdx-1]
		NumPop()
		var m rune
		m = Op[OpIdx-1]
		OpPop()
		var x float32
		if m == '+' {
			x = a + b
		}
		if m == '-' {
			x = a - b
		}
		if m == '*' {
			x = a * b
		}
		if m == '/' {
			x = a / b
		}
		NumPush(x)
	}
	return
}

func main() {
	pre := make(map[rune]int)
	pre['+'] = 1
	pre['-'] = 1
	pre['*'] = 2
	pre['/'] = 2

	var X string
	fmt.Scanln(&X)
	for i := 0; i < len(X); i++ {
		var c rune
		c = rune(X[i])
		var x float32
		if unicode.IsDigit(c) {
			x = 0
			for ; i < len(X) && unicode.IsDigit(rune(X[i])); i++ {
				num := float32(X[i] - '0')
				x = 10*x + num
			}
			if i < len(X) && X[i] == '.' {
				i++
				var AfterPoint float32
				AfterPoint = 0
				var Div float32
				Div = 10
				for ; i < len(X) && unicode.IsDigit(rune(X[i])); i++ {
					AfterPoint = AfterPoint + float32(X[i]-'0')/Div
					Div *= 10
				}
				x = x + AfterPoint
			}
			i--
			NumPush(x)
		} else if c == '(' {
			OpPush(c)
		} else if c == ')' {
			for Op[OpIdx-1] != '(' {
				eval()
			}
			OpPop()
		} else {
			for OpIdx > 0 && pre[Op[OpIdx-1]] >= pre[c] {
				eval()
			}
			OpPush(c)
		}
	}
	for OpIdx > 0 {
		eval()
	}
	fmt.Println(Num[0])
	return
}
