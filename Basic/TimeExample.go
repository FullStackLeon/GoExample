package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go func() {
		for {
			time.Sleep(3 * time.Second)
			ch1 <- "ch1"
		}

	}()

	go func() {
		for {
			time.Sleep(2 * time.Second)
			ch2 <- "ch2"
		}

	}()
	go func() {
		for {
			fmt.Println("Goroutine Num:", runtime.NumGoroutine())
			time.Sleep(5 * time.Second)
		}
	}()
	for {
		select {
		case data := <-ch1:
			fmt.Println("Received data from ch1:", data)
		case data := <-ch2:
			fmt.Println("Received data from ch2:", data)
		case <-time.After(2 * time.Second):
			fmt.Println("Timeout:", time.Now())
		}
	}
}
