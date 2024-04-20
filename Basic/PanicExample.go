package main

import "fmt"

func main() {
	// 无法捕获panicFunc中的Panic
	//go func() {
	//	defer func() {
	//		if err := recover(); err != nil {
	//			fmt.Println("main goroutine recover err", err)
	//		}
	//	}()
	//}()
	//
	// 无法捕获panicFunc中的Panic
	//defer func() {
	//	if err := recover(); err != nil {
	//		fmt.Println("main recover err", err)
	//	}
	//}()

	callPanicFunc()

	select {}
}

func panicFunc() {
	// 可以捕获panicFunc中的Panic
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panicFunc recover err", err)
		}
	}()
	panic("panicFunc for test")
}

func callPanicFunc() {
	// 无法捕获panicFunc中的Panic
	//defer func() {
	//	if err := recover(); err != nil {
	//		fmt.Println("callPanicFunc recover err", err)
	//	}
	//}()
	go panicFunc()
}
