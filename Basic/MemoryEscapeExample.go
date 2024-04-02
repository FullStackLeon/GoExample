package main

// 逃逸分析：go build -gcflags="-m" Basic/MemoryEscapeExample.go
import (
	"fmt"
)

var globalVar *int

// NewVar 在函数内创建变量x，将x的地址存储到全局变量中
func NewVar() {
	x := 10
	globalVar = &x
}

type Person struct {
	Name string
	Age  int
}

// NewPerson 局部变量p的指针返回到函数外部
func NewPerson(name string, age int) *Person {
	p := Person{Name: name, Age: age}
	return &p
}

func main() {
	// 场景1：局部变量指针返回到函数外部
	p := NewPerson("Jack", 18)
	fmt.Println(p)

	// 场景2：通过全局变量访问x的值，导致x逃逸到堆上
	NewVar()
	fmt.Println(*globalVar)
}
