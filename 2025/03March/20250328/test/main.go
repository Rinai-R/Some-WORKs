package main

import "fmt"

// 定义一个接口，包含方法 A 和 B
type IB interface {
	A()
	B()
}

// 父结构体
type Parent struct {
	IB // 嵌入接口
}

// 父类方法 A
func (p *Parent) A() {
	fmt.Println("Parent's A()")
	p.IB.B() // 调用接口的 B 方法，实现动态分派
}

// 父类默认的 B 方法
func (p *Parent) B() {
	fmt.Println("Parent's B()")
}

// 子结构体
type Child struct {
	Parent // 嵌入父结构体
}

// 子类重写 B 方法
func (c *Child) B() {
	fmt.Println("Child's B()")
}

func main() {
	child := &Child{}
	child.IB = child // 将接口绑定到子类实例
	child.A()        // 调用 A 方法，触发子类的 B 方法
}
