package main

import "fmt"

func producer(ch chan<- int) {
	defer close(ch)
	for i := 0; i < 10; i++ {
		ch <- i
	}
}

func consumer(ch <-chan int, done chan<- struct{}) {
	for v := range ch {
		fmt.Println("chan value:", v)
	}
	done <- struct{}{}
}

func main() {
	// 场景1：数据通信
	ch := make(chan int)
	done := make(chan struct{})
	go producer(ch)
	go consumer(ch, done)

	<-done
}
