package main

import "fmt"

type Chain struct {
	Name  string
	Chain *Chain
}

func (c *Chain) Handler() {
	fmt.Println(c.Name + "已处理，请求上一级处理")
	if c.Chain != nil {
		c.Chain.Handler()
	} else {
		fmt.Println("所有流程走完")
	}
}

func main() {
	fmt.Println("老板:实现逐级审批的场景")
	fmt.Println("你:好的老板")
	cT1 := &Chain{
		Name: "t1",
		Chain: &Chain{
			Name: "t2",
			Chain: &Chain{
				Name:  "t3",
				Chain: nil,
			},
		},
	}
	cT1.Handler()
}
