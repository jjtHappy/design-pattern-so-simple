package main

import "fmt"

type Product string

type AbstractFactory interface {
	createA() Product
	createB() Product
}

type Factory struct {
}

func (*Factory) createA() Product {
	fmt.Println("生产A")
	return Product("A")
}

func (*Factory) createB() Product {
	fmt.Println("生产B")
	return Product("B")
}

func main() {
	fmt.Println("从万能工厂那生产所有产品")
	fmt.Println("好的，老板")

	var f AbstractFactory
	f = &Factory{}
	f.createA()
	f.createB()
}
