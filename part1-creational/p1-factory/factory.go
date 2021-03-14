package main

import "fmt"

//工厂接口
type AbstractFactory interface {
	Produce() AbstractProduct
}

//抽象产品
type AbstractProduct string

//工厂实现
type AFactory struct {
}

func (ptr *AFactory) Produce() AbstractProduct {
	fmt.Println("工厂A开始生产产品A")
	return AbstractProduct("产品A") //产品实现
}

type BFactory struct {
}

func (ptr *BFactory) Produce() AbstractProduct {
	fmt.Println("工厂B开始生产产品B")
	return AbstractProduct("产品B")
}

func main() {
	fmt.Println("分别联系两个工厂A跟B生产产品A跟B")
	fmt.Println("好的老板")

	var factory AbstractFactory
	factory = &AFactory{}
	factory.Produce()
	factory = &BFactory{}
	factory.Produce()
}
