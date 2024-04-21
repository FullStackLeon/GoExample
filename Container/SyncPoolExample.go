package main

import (
	"fmt"
	"sync"
)

var slicePool = sync.Pool{
	New: func() interface{} {
		return make([]int, 0, 100)
	},
}

var mapPool = sync.Pool{
	New: func() interface{} {
		return make(map[int]int, 100)
	},
}

func main() {
	// slice sync pool 使用示例
	s1 := slicePool.Get().([]int)
	fmt.Printf("s1 pointer:%p\n", s1)
	s1 = append(s1, 1, 2, 3)
	fmt.Printf("s1 value:%d,s1 length:%d,s1 capacity:%d\n", s1, len(s1), cap(s1))
	s1 = s1[:0]
	slicePool.Put(s1)

	s2 := slicePool.Get().([]int)
	fmt.Printf("s2 pointer:%p\n", s2)
	fmt.Printf("s2 value:%d,s2 length:%d,s2 capacity:%d\n", s2, len(s2), cap(s2))

	// slice sync pool 输出示例：
	//s1 pointer:0xc00010e000
	//s1 value:[1 2 3],s1 length:3,s1 capacity:100
	//s2 pointer:0xc00010e000
	//s2 value:[],s2 length:0,s2 capacity:100

	// map sync pool 使用示例
	m1 := mapPool.Get().(map[int]int)
	fmt.Printf("m1 pointer:%p\n", m1)
	m1[1] = 1
	m1[2] = 2
	fmt.Printf("m1 value:%d,m1 length:%d\n", m1, len(m1))
	for k := range m1 {
		delete(m1, k)
	}
	mapPool.Put(m1)

	m2 := mapPool.Get().(map[int]int)
	fmt.Printf("m2 pointer:%p\n", m2)
	fmt.Printf("m2 value:%d,m2 length:%d\n", m2, len(m2))

	// map sync pool 输出示例
	// m1 pointer:0xc00007a0c0
	//m1 value:map[1:1 2:2],m1 length:2
	//m2 pointer:0xc00007a0c0
	//m2 value:map[],m2 length:0
}
