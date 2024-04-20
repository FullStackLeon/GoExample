package main

import "fmt"

func main() {
	var l1 = [2]string{
		"Hello", "World",
	}
	for _, v := range l1 {
		fmt.Println(v)
	}

	l2 := [2]string{"Hello", "World"}
	for _, v := range l2 {
		fmt.Println(v)
	}

	l3 := [5]string{1: "Hello", 3: "World"}
	for k, v := range l3 {
		fmt.Println(k, v)
	}

	l4 := [2][2]string{
		{"Hello", "World"},
		{"Hello", "World"},
	}
	for _, v := range l4 {
		fmt.Println(v)
	}

	l5 := [...]string{"Hello", "World"}
	for i := 0; i < len(l5); i++ {
		fmt.Println(l5[i])
	}

	var l6 [3]int
	fmt.Println(l6)
	l6[0] = 0
	l6[1] = 1
	l6[2] = 2
	for k, v := range l6 {
		fmt.Println(k, v)
	}

	fmt.Println(l6[1:])
}
