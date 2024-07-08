package Basic

import (
	"fmt"
	"testing"
)

func TestPanic(t *testing.T) {
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

	CallPanicFunc()

	select {}
}

func PanicFunc() {
	// 可以捕获panicFunc中的Panic
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("PanicFunc recover err", err)
		}
	}()
	panic("PanicFunc for test")
}

func CallPanicFunc() {
	// 无法捕获panicFunc中的Panic
	//defer func() {
	//	if err := recover(); err != nil {
	//		fmt.Println("callPanicFunc recover err", err)
	//	}
	//}()
	go PanicFunc()
}
