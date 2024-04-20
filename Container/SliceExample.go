package main

import (
	"fmt"
	"sort"
)

func main() {
	var s1 = []int{1, 2, 3}
	for _, v := range s1 {
		fmt.Println(v)
	}

	s2 := []int{1, 2, 3}
	for _, v := range s2 {
		fmt.Println(v)
	}

	s3 := make([]int, 3)
	s3[0] = 0
	s3[1] = 1
	s3[2] = 2
	for _, v := range s3 {
		fmt.Println(v)
	}

	s4 := make([]int, 3, 6)
	copy(s4, s3)
	s4 = append(s4, 3, 4, 5)
	for _, v := range s4 {
		fmt.Println(v)
	}

	l1 := [...]int{1, 2, 3, 4, 5}
	s5 := l1[1:]
	for _, v := range s5 {
		fmt.Println(v)
	}

	fmt.Println(len(s5))
	fmt.Println(cap(s5))

	// 切片排序
	s6 := make([]int, 0, 3)
	s6 = append(s6, 3, 2, 1)
	sort.Ints(s6)
	fmt.Println(s6)

	// 删除第二个元素
	s7 := append(s6[:1], s6[2:]...)
	fmt.Println(s7)

	// 切片反转
	s8 := []int{1, 2, 3}
	for i := 0; i < len(s8)/2; i++ {
		j := len(s8) - i - 1
		s8[i], s8[j] = s8[j], s8[i]
	}
	fmt.Println(s8)

	// 切片合并
	s9 := []int{1, 2, 3}
	s10 := []int{4, 5, 6}
	s11 := append(s9, s10...)
	fmt.Println(s11)
}
