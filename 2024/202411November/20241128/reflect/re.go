package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	age  int
}

func (stu Student) AGetName() {
	fmt.Println("调用AGetName方法")
	fmt.Println(stu.Name)
}

func (stu Student) BGetSum(a int, b int) int {
	fmt.Println("调用BGetSum方法")
	return a + b
}

func ReflectFunC(s interface{}) {
	val := reflect.ValueOf(s)
	fmt.Println(val)

	val.Method(0).Call(nil)

	var params []reflect.Value
	params = append(params, reflect.ValueOf(120))
	params = append(params, reflect.ValueOf(111))
	x1 := val.Method(1).Call(params)
	fmt.Println(x1[0].Int())
	val.Elem().Field(0).SetString("长江黄河")
}

func main() {
	s := Student{
		Name: "Rinai_",
		age:  15,
	}

	ReflectFunC(&s)
	fmt.Println(s)
	return
}
