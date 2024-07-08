package Basic

import (
	"fmt"
	"testing"
	"time"
)

func TestGoroutine(t *testing.T) {
	go func() {
		go func() {
			defer func() {
				if err := recover(); err != nil {
					fmt.Println("异常被捕获")
				}
			}()
			panic("子协程Panic")
		}()

		fmt.Println("父协程执行")
	}()

	time.Sleep(time.Second)
}

func TestSubGoroutineNoRecover(t *testing.T) {
	go func() {
		go func() {
			panic("子协程Panic")
		}()

		fmt.Println("父协程执行")
	}()

	time.Sleep(time.Second)
}
