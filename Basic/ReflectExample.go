package main

import (
	"fmt"
	"reflect"
)

type language struct {
	Name string
	Age  int
}

func main() {
	p := language{Name: "Golang", Age: 15}
	t := reflect.TypeOf(p)
	fmt.Printf("Type:%v#\n", t)
	fmt.Printf("Name:%v\n", t.Name())
	fmt.Printf("Kind:%v\n", t.Kind())
	fmt.Printf("NumFiled:%v\n", t.NumField())
	for i := 0; i < t.NumField(); i++ {
		fmt.Printf("Field Name:%v,Filed Type:%v\n", t.Field(i).Name, t.Field(i).Type)
	}
	v := reflect.ValueOf(p)
	fmt.Printf("Value:%v\n", v)
	fmt.Printf("Kind:%v\n", v.Kind())
	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("Field Type:%v,Filed Value:%v\n", v.Field(i).Type(), v.Field(i).Interface())
	}
}
