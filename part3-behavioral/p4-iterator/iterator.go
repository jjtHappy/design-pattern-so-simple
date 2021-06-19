package main

import "fmt"

type Iterator interface {
	HasNext() bool
	Next() string
	SetList([]string)
}

type IteratorImpl struct {
	index int32
	list  []string
}

func (i *IteratorImpl) HasNext() bool {
	return i.index < int32(len(i.list))
}

func (i *IteratorImpl) Next() string {
	r := i.list[i.index]
	i.index++
	return r
}

func (i *IteratorImpl) SetList(list []string) {
	i.list = list
}

type ListOwner struct {
	List []string
}

func (io *ListOwner) Iterator() (i Iterator) {
	i = &IteratorImpl{}
	i.SetList(io.List)
	return i
}

func main() {
	fmt.Println("老板：实现一个迭代器")
	fmt.Println("你：好的老板")
	lo := ListOwner{
		List: []string{"hello", "world"},
	}
	i := lo.Iterator()
	for i.HasNext() {
		fmt.Println(i.Next())
	}
}
