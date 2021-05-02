package main

import "fmt"

type Concrete struct {
}

func (c *Concrete) foo1() string {
	return "复杂的操作一"
}

func (c *Concrete) foo2() string {
	return "复杂的操作二"
}

func (c *Concrete) foo3() string {
	return "复杂的操作三"
}

type Facade struct {
	Concrete Concrete
}

func (f *Facade) foo() string {
	fmt.Println(f.Concrete.foo1())
	fmt.Println(f.Concrete.foo2())
	fmt.Println(f.Concrete.foo3())
	return "done"
}

func main() {
	fmt.Println("想个办法封装一个复杂的流程")
	fmt.Println("好的")
	f := &Facade{
		Concrete: Concrete{},
	}
	f.foo()
}
