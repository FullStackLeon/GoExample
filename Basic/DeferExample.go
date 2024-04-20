package main

import (
	"fmt"
)

var gVar = 100

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Panic err", err)
		}
	}()
	lVar := 200
	defer func(v int) {
		fmt.Printf("defer in main function,gVar:%v,v:%v\n", gVar, v)
	}(lVar)
	defer func() {
		fmt.Printf("defer in main function,gVar:%v,lVar:%v\n", gVar, lVar)
	}()
	gVar++
	lVar++
	returnFunc()
	panicFunc()
	// 输出结果：
	//defer in return example,gVar:101
	//defer in panic example,gVar:102
	//defer in main function,gVar:103,lVar:201
	//defer in main function,gVar:103,v:200
	//Panic err panic for test
}

func returnFunc() {
	defer fmt.Printf("defer in return example,gVar:%v\n", gVar)
	gVar++
	return
}

func panicFunc() {
	defer fmt.Printf("defer in panic example,gVar:%v\n", gVar)
	gVar++
	panic("panic for test")
}
