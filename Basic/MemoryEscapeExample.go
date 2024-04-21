package main

// 逃逸分析：go build -gcflags="-m" Basic/MemoryEscapeExample.go
//-m 参数在逃逸分析中还有其他取值，主要用于调试和优化。常见取值包括：
//-m=0：关闭逃逸分析，不进行逃逸分析，也不输出逃逸分析的信息。
//-m=1：开启逃逸分析，并输出详细的逃逸分析信息。
//-m=2：开启逃逸分析，但只输出逃逸分析的结果，不显示详细信息。
import (
	"fmt"
)

var globalVar *int

// NewVar 在函数内创建局部变量x，将局部变量x的地址存储到全局变量中
func NewVar() {
	x := 10 // x 逃逸到堆上
	globalVar = &x
}

type Person struct {
	Name string
	Age  int
}

// NewPerson 局部变量p的指针返回到函数外部
func NewPerson(name string, age int) *Person {
	p := Person{Name: name, Age: age} // p 逃逸到堆上
	return &p
}

func bar() {
	x := 20
	foo(&x) // 局部变量x 逃逸到堆上
}

func foo(x *int) {
	fmt.Println(x) // x二次逃逸到堆上
}

func closure() func() {
	x := 30
	return func() {
		fmt.Println(x) // 局部变量x 逃逸到堆上
	}
}
func main() {
	// 场景1：局部变量指针或引用类型返回到函数外部
	p := NewPerson("Jack", 18)
	fmt.Println(p)

	// 场景2：通过全局变量访问局部变量x的值，导致局部变量x逃逸到堆上
	NewVar()
	fmt.Println(*globalVar)
	// 场景3：将局部变量x地址作为函数参数传入方法
	bar()

	// 场景4：闭包函数中引用局部变量时
	closure()() // 闭包函数中的x二次逃逸到堆上
}
