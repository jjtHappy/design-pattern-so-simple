package main

import "fmt"

type Mediator interface {
	Notify()
	Register(*Member)
}

type MediatorImpl struct {
	members []*Member
}

func (m *MediatorImpl) Notify() {
	for _, m := range m.members {
		m.Notify()
	}
}

func (m *MediatorImpl) Register(member *Member) {
	m.members = append(m.members, member)
}

type Member struct {
	Name string
}

func (m *Member) Notify() {
	fmt.Println(m.Name + "收到通知")
}

func main() {
	mediator := &MediatorImpl{}
	mediator.Register(&Member{
		Name: "成员1",
	})
	mediator.Register(&Member{
		Name: "成员2",
	})
	mediator.Register(&Member{
		Name: "成员3",
	})
	mediator.Notify()
}
