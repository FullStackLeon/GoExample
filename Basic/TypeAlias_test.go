package Basic

import (
	"fmt"
	"testing"
)

type IInterface1 interface {
	method1()
}

type User struct {
}

type MyUser = User

func (i User) method() {
	fmt.Println("User.method")
}

func (i MyUser) method1() {
	fmt.Println("MyUser.method1")
}

// type alias 场景1：MyUser实现了IInterface1接口，等同于User结构体也实现了
func Test1(t *testing.T) {
	var u User
	var u1 MyUser
	var i IInterface1 = u
	var i1 IInterface1 = u1

	// u的类型是User，MyUser是User的别名，u1实例也有方法method
	// MyUser隐式实现了IInterface1接口，那么User同样实现了IInterface1
	// 将u和u1 均可复制给IInterface1接口类型的i和i1变量，i和i1均有方法method1
	u.method()
	u1.method()
	i.method1()
	i1.method1()
}

type IInterface2 interface {
	method2()
}

type MyInterface1 IInterface2
type MyInterface2 = IInterface2

type MyInt int

func (i MyInt) method2() {
	fmt.Println("MyInt.method2")
}

// MyInterface1 基于 IInterface2定义的新类型
// MyInterface2 是IInterface2的别名
// type alias 场景2：不论是新类型还是别名，只要MyInt实现了方法method2，IInterface2，MyInterface1和MyInterface2 三个类型间可以互相赋值
func Test2(t *testing.T) {
	var i IInterface2 = MyInt(100)
	var i1 MyInterface1 = MyInt(100)
	var i2 MyInterface2 = MyInt(100)

	i = i1
	i = i2
	i1 = i2
	i1 = i
	i2 = i
	i2 = i1
}
