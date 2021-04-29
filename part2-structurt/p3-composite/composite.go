package main

import "fmt"

type Composite struct {
	Name    string
	Childes []*Composite
}

func (c *Composite) AddChildes(child *Composite) {
	if c.Childes == nil {
		c.Childes = []*Composite{}
	}
	c.Childes = append(c.Childes, child)
}

func main() {
	fmt.Println("实现一个组织架构")
	fmt.Println("好的老板")
	p := &Composite{
		Name: "公司",
	}
	p.AddChildes(&Composite{
		Name: "总办",
	})
	p.AddChildes(&Composite{
		Name: "开发部",
		Childes: []*Composite{
			&Composite{
				Name: "开发一部",
			},
			&Composite{
				Name: "开发二部",
			},
		},
	})
	printComposite(p, 1)
}

func printComposite(p *Composite, offset int) {
	for i := 0; i < offset; i++ {
		fmt.Print("-")
	}
	fmt.Println(p.Name)
	for _, c := range p.Childes {
		printComposite(c, offset+1)
	}
}
