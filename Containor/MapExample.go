package main

import (
	"fmt"
)

func main() {
	m := make(map[string]interface{}, 10)
	m["Golang"] = struct{ name string }{name: "Golang"}
	m["Java"] = struct{ name string }{name: "Java"}
	m["Rust"] = struct{ name string }{name: "Rust"}
	m["Python"] = struct{ name string }{name: "Python"}
	for k, v := range m {
		fmt.Println(k, v)
	}

	delete(m, "Golang")
	m["Java"] = struct{ name string }{name: "Java SpringBoot"}
	for k, v := range m {
		fmt.Printf("key=%v,value=%#v\n", k, v)
	}
}
