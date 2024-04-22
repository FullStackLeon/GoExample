package main

import "fmt"

type Animal struct {
	Name string
}

// ChangeNameValue 修改的是值接收者的副本，不会影响原始对象的name
func (a Animal) ChangeNameValue(newName string) {
	a.Name = newName
}

// ChangeNamePointer 修改的指针接收者的副本，会影响原始对象的name
func (a *Animal) ChangeNamePointer(newName string) {
	a.Name = newName
}

func main() {
	a := Animal{Name: "Cat"}

	a.ChangeNameValue("Dog")
	fmt.Println("Name after using value receiver:", a.Name) // 输出：Cat

	a.ChangeNamePointer("Charlie")
	fmt.Println("Name after using pointer receiver:", a.Name) // 输出：Dog
}
