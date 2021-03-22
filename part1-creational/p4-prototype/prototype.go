package main

import "fmt"

type AbstractPrototype interface {
	clone() AbstractPrototype
}

type Prototype struct {
	field1 string
	field2 string
	field3 string
}

func (o Prototype) clone() AbstractPrototype {
	return o
}

func main() {
	fmt.Println("实现一个可复制的对象")
	fmt.Println("好的老板")
	p := &Prototype{
		field1: "1",
		field2: "2",
		field3: "3",
	}
	fmt.Printf("%+v", p.clone())
}
