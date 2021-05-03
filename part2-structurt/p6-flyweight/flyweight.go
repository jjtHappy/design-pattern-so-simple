package main

import "fmt"

type flyweight struct {
}

func (f *flyweight) Name() string {
	return "这是系统中唯一的蝇量对象"
}

type FlyweightFactory struct {
	flyweight *flyweight
}

func (receiver *FlyweightFactory) Get() *flyweight {
	if receiver.flyweight == nil {
		receiver.flyweight = &flyweight{}
	}
	return receiver.flyweight
}

func main() {
	fmt.Println("实现一个不可变对象的复用")
	fmt.Println("好的老板")
	fact := &FlyweightFactory{}
	fmt.Println(fact.Get().Name())
}
