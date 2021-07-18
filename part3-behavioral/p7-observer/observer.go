package main

import "fmt"

type Observer interface {
	Notify(t *TargetSubject)
}

type ObserverImpl struct {
}

func (*ObserverImpl) Notify(t *TargetSubject) {
	fmt.Printf("观察者监听到状态变换:%+v", t.State)
}

type TargetSubject struct {
	State     bool
	Observers []Observer
}

func (t *TargetSubject) ChangeStatue() {
	t.State = true
	fmt.Println("目标对象状态被改变:", t.State)
	for _, o := range t.Observers {
		o.Notify(t)
	}
}

func main() {
	t := &TargetSubject{
		State: false,
		Observers: []Observer{
			&ObserverImpl{},
		},
	}
	t.ChangeStatue()
}
