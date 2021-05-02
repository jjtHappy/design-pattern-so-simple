package main

import "fmt"

type Interface interface {
	foo() string
}

type ConcreteImpl struct {
}

func (c ConcreteImpl) foo() string {
	return "真正的实现在干活"
}

type Decorator struct {
	ConcreteImpl Interface
}

func (d Decorator) foo() string {
	return "装饰者增加了功能，再让" + d.ConcreteImpl.foo()
}

func main() {
	fmt.Println("在不动原有代码的情况下，给现有的实现类增加功能")
	fmt.Println("好的，老板")
	d := &Decorator{
		ConcreteImpl: &ConcreteImpl{},
	}
	fmt.Println(d.foo())
}
