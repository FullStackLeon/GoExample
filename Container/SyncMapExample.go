package main

import (
	"fmt"
	"sync"
)

var me sync.Map

func main() {
	me.Store("apple", 10)
	fmt.Println(me.Load("apple"))

	prevPrice, loaded := me.Swap("apple", 12)
	fmt.Println(prevPrice, loaded)
	fmt.Println(me.Load("apple"))

	// actual: 如果键存在于映射中，actual 是键对应的旧值；如果键不存在，actual 是存储的新值。
	// loaded: 一个布尔值，表示键是否存在于映射中：
	// 如果 loaded == true，说明键存在，actual 是获取的旧值。
	// 如果 loaded == false，说明键不存在，actual 是存储的新值。
	pearPreValue, loaded := me.LoadOrStore("pear", 5)
	fmt.Println(pearPreValue, loaded)

	pearPreValue, loaded = me.LoadOrStore("pear", 6)
	fmt.Println(pearPreValue, loaded)

	me.LoadOrStore("orange", 3)
	fmt.Println("Range over the map:")
	me.Range(func(key, value any) bool {
		fmt.Println(key, value)
		return true
	})

	fmt.Println("CompareAndSwap:")
	fmt.Println(me.CompareAndSwap("apple", 12, 10))
	fmt.Println(me.Load("apple"))
	fmt.Println("CompareAndDelete:")
	fmt.Println(me.CompareAndDelete("apple", 10))
	fmt.Println(me.Load("apple"))

	// Output:
	// 10 true
	// 10 true
	// 12 true
	// 5 false
	// 5 true
	// Range over the map:
	// apple 12
	// pear 5
	// orange 3
	// CompareAndSwap:
	// true
	// 10 true
	// CompareAndDelete:
	// true
	// <nil> false
}
