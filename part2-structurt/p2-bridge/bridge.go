package main

import "fmt"

type Interface interface {
	foo() string
}

type InterfaceImplA struct {
}

type InterfaceImplB struct {
}

func (*InterfaceImplA) foo() string {
	return "实现A在干活"
}

func (*InterfaceImplB) foo() string {
	return "实现B在干活"
}

type Bridge struct {
	worker Interface
}

func (m *Bridge) Init(p Interface) {
	m.worker = p
}

func (m *Bridge) DoSomething() string {
	return m.worker.foo()
}

func main() {
	fmt.Println("告诉我啥叫tm的面向接口编程")
	fmt.Println("这就是tm的面向接口编程")
	b := &Bridge{}
	b.Init(&InterfaceImplA{})
	fmt.Println(b.DoSomething())
	b.Init(&InterfaceImplB{})
	fmt.Println(b.DoSomething())
}
