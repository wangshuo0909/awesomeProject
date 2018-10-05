package main

import "fmt"

type(
	Stackk struct {
		top *node
		lenth int
	}
	node struct {
		value interface{}
		prew *node
	}

)

func NewStackk() *Stackk{
	return &Stackk{
		nil,
		0,
	}
}
func (s *Stackk)Push(value interface{})  {
	s.lenth++
	node:=&node{value,s.top}
	s.top=node
}
func (s *Stackk)Pop() interface{} {
	if s.lenth==0 {
		panic("empty")
	}
	s.lenth--
	popV:=s.top
	s.top=s.top.prew
	return popV.value
}
func (s *Stackk)Peek()interface{}  {
	if s.lenth==0 {
		panic("empty")
	}
	peek:=s.top
	return peek
}
func main()  {
	stack:=NewStackk()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}