package Basic

import (
	"fmt"
	"testing"
)

var GVar = 100

func TestDefer(t *testing.T) {
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
