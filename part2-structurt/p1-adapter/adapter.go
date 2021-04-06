package main

import "fmt"

type Interface interface {
	foo() string
}

type RealObject struct {
}

func (p *RealObject) Do() string {
	return "真实对象在做事"
}

type Adapter struct {
	realObject *RealObject
}

func (p *Adapter) init(object *RealObject) {
	p.realObject = object
}

func (p *Adapter) foo() string {
	if p.realObject == nil {
		fmt.Println("你需要初始化一个对象")
	}
	return "适配器让" + p.realObject.Do()
}

func main() {
	fmt.Println("设计一个接口适配器，让任意真实对象可以利用这个适配器适配接口")
	fmt.Println("好的老板")
	realObject := &RealObject{}
	adapter := &Adapter{}
	adapter.init(realObject)
	fmt.Println(adapter.foo())
}
