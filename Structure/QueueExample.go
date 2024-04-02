package main

import "fmt"

type Queue struct {
	items []interface{}
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(item interface{}) {
	q.items = append(q.items, item)
}
func (q *Queue) Poll() interface{} {
	if len(q.items) == 0 {
		return nil
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func (q *Queue) Peek() interface{} {
	if len(q.items) == 0 {
		return nil
	}
	return q.items[0]
}

func (q *Queue) Len() int {
	return len(q.items)
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func main() {
	s := NewQueue()
	s.Enqueue(1)
	s.Enqueue(2)
	s.Enqueue(3)
	fmt.Println(s.Peek(), s.Len(), s.IsEmpty())
	s.Poll()
	fmt.Println(s.Peek(), s.Len(), s.IsEmpty())
	s.Poll()
	s.Poll()
	fmt.Println(s.Peek(), s.Len(), s.IsEmpty())
}
