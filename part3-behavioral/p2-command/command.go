package main

import "fmt"

type Command interface {
	Exec()
}

type Server struct {
}

func (*Server) Foo1() {
	fmt.Println("方法一")
}
func (*Server) Foo2() {
	fmt.Println("方法二")
}

type CommandImpl1 struct {
	Server *Server
}

func (c *CommandImpl1) Exec() {
	fmt.Println("命令一执行")
	c.Server.Foo1()
	c.Server.Foo2()
}

type CommandImpl2 struct {
	Server *Server
}

func (c *CommandImpl2) Exec() {
	fmt.Println("命令二执行")
	c.Server.Foo2()
	c.Server.Foo1()
}

func main() {
	fmt.Println("老板：对一个业务实现不同行为的封装")
	fmt.Println("你：老板")
	server := &Server{}
	command1 := CommandImpl1{
		Server: server,
	}
	command1.Exec()
	command2 := CommandImpl1{
		Server: server,
	}
	command2.Exec()
}
