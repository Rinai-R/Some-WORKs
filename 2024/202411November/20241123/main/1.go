package main

import (
	"20241123/student"
	"fmt"
)

func main() {
	st := student.Student{"Rinai", 19}
	fmt.Println(st)
	return
}

//package main
//
//import (
//	"20241123/student"
//	"fmt"
//)
//
//func (s *student.Student) String() string {
//	st := fmt.Sprintf("名叫%v, 今年%v岁", (*s).Name, (*s).age)
//	return st
//}
//
//func main() {
//	std := student.Student{
//		"Rinai",
//		19,
//	}
//	fmt.Println(&std)
//	return
//}
