package main

import (
	"fmt"
)

var GVar = 100

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Panic err", err)
		}
	}()
	lVar := 200
	defer func(v int) {
		fmt.Printf("defer in main function,GVar:%v,v:%v\n", GVar, v)
	}(lVar)
	defer func() {
		fmt.Printf("defer in main function,GVar:%v,lVar:%v\n", GVar, lVar)
	}()
	GVar++
	lVar++
	ReturnFunc()
	DeferPanicFunc()
	// 输出结果：
	//defer in return example,GVar:101
	//defer in panic example,GVar:102
	//defer in main function,GVar:103,lVar:201
	//defer in main function,GVar:103,v:200
	//Panic err panic for test
}

func ReturnFunc() {
	defer fmt.Printf("defer in return example,GVar:%v\n", GVar)
	GVar++
	return
}

func DeferPanicFunc() {
	defer fmt.Printf("defer in panic example,GVar:%v\n", GVar)
	GVar++
	panic("panic for test")
}
