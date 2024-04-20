package main

import (
	"fmt"
	"strings"
)

func main() {
	strList := []string{
		"Hello", "World", "Golang",
	}

	// 方式1：strings.Join
	fmt.Println(strings.Join(strList, ","))
	// 方式2：strings.Builder
	var builder strings.Builder
	length := len(strList)
	for index, s := range strList {
		builder.WriteString(s)
		if index != length-1 {
			builder.WriteString(",")
		}
	}
	fmt.Println(builder.String())
}
