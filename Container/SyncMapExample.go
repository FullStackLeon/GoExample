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

	me.LoadOrStore("pear", 5)
	me.LoadOrStore("orange", 3)
	me.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})

	fmt.Println(me.CompareAndSwap("apple", 12, 10))
	fmt.Println(me.Load("apple"))
	fmt.Println(me.CompareAndDelete("apple", 10))
	fmt.Println(me.Load("apple"))
}
