package main

import "fmt"

type Config struct {
	Name string
	Q    chan string
	In   interface{}
}

func main() {
	hello := interface{}("hello")
	fmt.Println(hello)

	i := interface{}(1)
	fmt.Println(i)

	e1 := interface{}(nil)
	fmt.Println(e1)

	var e2 interface{}
	if e1 == e2 {
		fmt.Println("Equal e1 == e2")
	} else {
		fmt.Println("Not equal e1 == e2")
	}
	q := make(chan string, 100)
	c1 := Config{
		Name: "Jack",
		Q:    q,
		In:   interface{}("hello"),
	}
	c2 := Config{
		Name: "Jack",
		Q:    q,
		In:   "hello",
	}

	if c1 == c2 {
		fmt.Println("Equal c1 == c2")
	} else {
		fmt.Println("Not equal c1 == c2")
	}
}
