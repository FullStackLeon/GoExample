package main

import (
	"fmt"
	"sync/atomic"
	"unsafe"
)

func main() {
	// atomic Pointer Generics
	fmt.Println("---- atomic.Pointer Generics ----")
	var ps1 atomic.Pointer[string]
	str1 := "Hello, world!"
	ps1.Store(&str1)

	fmt.Println(*ps1.Load())

	str2 := "Goodbye, world!"
	fmt.Println(*ps1.Swap(&str2))
	fmt.Println(*ps1.Load())

	fmt.Println("compare and swap:", ps1.CompareAndSwap(&str2, &str1))
	fmt.Println(*ps1.Load())

	// atomic Uintptr
	fmt.Println("---- atomic.Uintptr ----")
	u1 := uintptr(unsafe.Pointer(&str1))
	fmt.Println(u1)

	var au1 atomic.Uintptr
	au1.Store(u1)
	fmt.Println(au1.Load())

	u2 := uintptr(unsafe.Pointer(&str2))
	fmt.Println(au1.Swap(u2))
	fmt.Println(au1.Load())

	fmt.Println("compare and swap:", au1.CompareAndSwap(u2, u1))
	fmt.Println(au1.Load())

	// 普通指针转为uintptr
	fmt.Println("---- uintptr & Pointer ----")
	u3 := uintptr(unsafe.Pointer(&str1))
	fmt.Println(u3)

	// uintptr转为普通指针
	var str3 *string = (*string)(unsafe.Pointer(u3))
	fmt.Println(*str3)

	// ---- atomic.Pointer Generics ----
	// Hello, world!
	// Hello, world!
	// Goodbye, world!
	// compare and swap: true
	// Hello, world!
	// ---- atomic.Uintptr ----
	// 824635293984
	// 824635293984
	// 824635293984
	// 824635294016
	// compare and swap: true
	// 824635293984
	// ---- uintptr & Pointer ----
	// 824635293984
	// Hello, world!
}
