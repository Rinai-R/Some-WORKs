package main

//	func main() {
//		var f = []int{1, 2, 3}
//		//range只会遍历初始的长度，而len不会！
//		for i := 0; i < len(f); i++ {
//			f = append(f, f[i])
//			fmt.Println("v=", f[i])
//		}
//		return
//	}
func main() {
	hash := map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
	}

	for k, v := range hash {
		println(k, v)
	}
}
