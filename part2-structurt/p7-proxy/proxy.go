package main

import "fmt"

type Concrete struct {
}

func (c *Concrete) Foo() string {
	return "真正的对象在干活"
}

type Proxy struct {
	Concrete *Concrete
}

func (p *Proxy) Foo() string {
	fmt.Println("代理在对象干活前偷偷的做了某些事情")
	fmt.Println(p.Concrete.Foo())
	fmt.Println("代理在对象干活后偷偷的做了某些事情")
	return "代理帮你干了这个活"
}

func main() {
	fmt.Println("在不改变某个类的动作下，在他调用的前后注入一些其它事件")
	fmt.Println("好的，老板")
	p := &Proxy{
		Concrete: &Concrete{},
	}
	fmt.Println(p.Foo())
}
