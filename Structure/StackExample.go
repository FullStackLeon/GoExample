package main

import "fmt"

type Stack struct {
	items []interface{}
}

func NewStack() *Stack {
	return &Stack{}
}
func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() interface{} {
	if len(s.items) == 0 {
		return nil
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func (s *Stack) Peek() interface{} {
	if len(s.items) == 0 {
		return nil
	}
	return s.items[len(s.items)-1]
}
func (s *Stack) Len() int {
	return len(s.items)
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Clear() {
	s.items = s.items[:0]
}

func main() {
	s := NewStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(s.Peek(), s.Len(), s.IsEmpty())
	s.Pop()
	fmt.Println(s.Peek(), s.Len(), s.IsEmpty())
	s.Pop()
	s.Pop()
	fmt.Println(s.Peek(), s.Len(), s.IsEmpty())
}
