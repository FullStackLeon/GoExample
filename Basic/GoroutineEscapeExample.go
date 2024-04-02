package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

//func worker() {
//	resp, err := http.Get("https://go.dev/")
//	if err != nil {
//		fmt.Println("Error:", err)
//		return
//	}
//	_, _ = io.ReadAll(resp.Body)
//	//err = resp.Body.Close()
//	//if err != nil {
//	//	fmt.Println("Error closing body:", err)
//	//}
//}

func wgWorker(wg *sync.WaitGroup) {
	defer wg.Done()
	// Do some work...
}

func timerWorker() {
	timer := time.NewTicker(1 * time.Second)
	for {
		t := <-timer.C
		fmt.Println("Tick:", t)
	}
}

func chWorker(ch chan int) {
	for {
		select {
		case _, ok := <-ch:
			if !ok {
				return
			}
		default:
		}

	}
}

func main() {
	// 场景1：未关闭resp.Body.Close() ???
	//for index := 0; index < 5; index++ {
	//	go worker()
	//}
	//time.Sleep(1000 * time.Millisecond)
	//fmt.Printf("场景1:goroutine数量 %d\n", runtime.NumGoroutine())

	// 场景2：忘记调用wg.Wait()
	var wg2 sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg2.Add(1)
		go wgWorker(&wg2)
	}
	// wg.Wait()
	fmt.Printf("场景2:goroutine数量 %d\n", runtime.NumGoroutine())
	//
	//// 场景3：未停止定时器，导致Goroutine泄露
	go timerWorker()
	fmt.Printf("场景3:goroutine数量 %d\n", runtime.NumGoroutine())

	// 场景4：未关闭ch，导致Goroutine泄露
	ch := make(chan int)
	go chWorker(ch)

	ch <- 1
	ch <- 2
	// close(ch)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("场景4:goroutine数量 %d\n", runtime.NumGoroutine())
}
