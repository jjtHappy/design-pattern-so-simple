package main

import "fmt"

type Memento struct {
	State int32
}

type Originator struct {
	State int32
}

func (o *Originator) CreateMemento() *Memento {
	fmt.Println("创建备忘录")
	return &Memento{
		State: o.State,
	}
}

func (o *Originator) RestoreFromMemento(memento *Memento) {
	fmt.Println("回滚")
	o.State = memento.State
}

func main() {
	fmt.Println("老板：对我们的现有记录做存档，后面可以恢复")
	fmt.Println("你：好的老板")
	o := &Originator{State: 100}
	fmt.Printf("记录状态[%+v]\n", o.State)
	m := o.CreateMemento()
	o.State -= 1
	fmt.Printf("发生变动，记录状态[%+v]\n", o.State)
	o.RestoreFromMemento(m)
	fmt.Printf("记录状态[%+v]\n", o.State)
}
