package main

import "fmt"

type Product struct {
	FieldA string
	FieldB string
	FieldC string
	FieldD string
}

type AbstractBuilder interface {
	Init()
	SetFieldA()
	SetFieldB()
	SetFieldC()
	SetFieldD()
	Build() Product
}

type BuilderA struct {
	product *Product
}

func (b *BuilderA) Init() {
	b.product = &Product{}
}

func (b *BuilderA) SetFieldA() {
	if b.product == nil {
		panic("need init")
	}
	b.product.FieldA = "A"
}

func (b *BuilderA) SetFieldB() {
	if b.product == nil {
		panic("need init")
	}
	b.product.FieldB = "B"
}

func (b *BuilderA) SetFieldC() {
	if b.product == nil {
		panic("need init")
	}
	b.product.FieldC = "C"
}

func (b *BuilderA) SetFieldD() {
	if b.product == nil {
		panic("need init")
	}
	b.product.FieldD = "D"
}

func (b *BuilderA) Build() Product {
	return *b.product
}

func main() {
	fmt.Println("指导建造者搞一个复杂的产品")
	fmt.Println("好的")
	var builder AbstractBuilder
	builder = &BuilderA{}
	builder.Init()
	builder.SetFieldA()
	builder.SetFieldB()
	builder.SetFieldC()
	builder.SetFieldD()
	fmt.Printf("产品[%+v]", builder.Build())
}
