package main

import (
	"fmt"
	"sync"
)

type Singleton struct {
	Name string
}

var singletonHungry = &Singleton{
	Name: "singletonHungry",
}
var singletonLazy *Singleton
var once sync.Once

func GetSingletonHungry() *Singleton {
	return singletonHungry
}

func GetSingletonLazy() *Singleton {
	if singletonLazy == nil {
		once.Do(func() {
			singletonLazy = &Singleton{
				Name: "singletonLazy",
			}
		})
	}
	return singletonLazy
}

func main() {
	fmt.Println("设计一个单例模式")
	fmt.Println("好的老板")
	fmt.Println(GetSingletonHungry().Name)
	fmt.Println(GetSingletonLazy().Name)
}
