package main

type Stack struct {
	arr []int
	top int
}

func NewStack() *Stack {
	return &Stack{
		arr: make([]int, 64),
		top: -1,
	}
}
func (s *Stack) push(i int) {
	s.top++
	s.arr[s.top] = i
}

func (s *Stack) pop() (int) {
	if s.top < 0 {
		panic("empty")
	}
	v := s.arr[s.top]
	s.top--
	return v
}

func main() {
	s := NewStack()
	s.push(1)
	s.push(2)
	s.push(3)
	println(s.pop())
	println(s.pop())
	println(s.pop())
	println(s.pop())
}
