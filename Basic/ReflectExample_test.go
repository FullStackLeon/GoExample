package Basic

import (
	"fmt"
	"reflect"
	"testing"
)

type language struct {
	Name string
	Age  int
}

func TestReflect(t *testing.T) {
	p := language{Name: "Golang", Age: 15}
	typ := reflect.TypeOf(p)
	fmt.Printf("Type:%v#\n", typ)
	fmt.Printf("Name:%v\n", typ.Name())
	fmt.Printf("Kind:%v\n", typ.Kind())
	fmt.Printf("NumFiled:%v\n", typ.NumField())
	for i := 0; i < typ.NumField(); i++ {
		fmt.Printf("Field Name:%v,Filed Type:%v\n", typ.Field(i).Name, typ.Field(i).Type)
	}
	v := reflect.ValueOf(p)
	fmt.Printf("Value:%v\n", v)
	fmt.Printf("Kind:%v\n", v.Kind())
	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("Field Type:%v,Filed Value:%v\n", v.Field(i).Type(), v.Field(i).Interface())
	}
}
