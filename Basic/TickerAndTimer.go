package main

import (
	"fmt"
	"time"
)

func main() {
	// NewTicker 示例
	ticker := time.NewTicker(100 * time.Millisecond)
	go func() {
		for v := range ticker.C {
			fmt.Println("Ticker value: ", v)
		}
	}()
	time.Sleep(time.Second)
	ticker.Stop()

	// NewTimer 示例
	timer := time.NewTimer(100 * time.Millisecond)
	go func() {
		for v := range timer.C {
			fmt.Println("Timer value: ", v)
		}
	}()
	time.Sleep(time.Second)
	timer.Stop()
}
