package Basic

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"runtime"
	"testing"
	"time"
)

func TestGoroutine(t *testing.T) {
	go func() {
		go func() {
			// 协成是独立存在，不存在父子关系，如果内部协成Panic未被捕获，外部协成将终止
			defer func() {
				if err := recover(); err != nil {
					fmt.Printf("异常被捕获,err:%v\n", err)
				}
			}()
			panic(errors.New("内部协程Panic"))
		}()

		fmt.Println("外部协程执行")
	}()

	time.Sleep(time.Second)
	// Output:
	//外部协程执行
	//异常被捕获,err:内部协程Panic
}

func TestSubGoroutineNoRecover(t *testing.T) {
	go func() {
		go func() {
			panic("内部协程Panic")
		}()

		fmt.Println("外部协程执行")
	}()

	time.Sleep(time.Second)
}

func TestGoroutineNum(t *testing.T) {
	//Go 的 testing 包在运行测试时，会自动创建一个协程来管理测试案例的执行。在测试环境中，通常会看到 2 个初始协程：
	//一个主协程:Go测试工具会在后台生成一个专用的 main 函数，负责加载和执行测试案例
	//一个测试管理协程:负责运行测试案例，并生成测试结果。
	fmt.Println("before goroutine num:", runtime.NumGoroutine())

	for i := 0; i < 10; i++ {
		go func() {
			time.Sleep(time.Second)
		}()
	}
	time.Sleep(time.Millisecond * 100)
	fmt.Println("after goroutine num:", runtime.NumGoroutine())
	// Output:
	//before goroutine num: 2
	//after goroutine num: 12
}

func TestGoroutineMem(t *testing.T) {
	readMemStats()
	for i := 0; i < 10; i++ {
		go func() {
			time.Sleep(time.Second)
		}()
	}

	time.Sleep(time.Millisecond * 100)
	readMemStats()
	// Output:
	//Alloc = 204 KB	TotalAlloc = 204 KB	Sys = 8209 KB	HeapAlloc = 204 KB	HeapSys = 3808 KB	HeapIdle = 3032 KB
	//Alloc = 210 KB	TotalAlloc = 210 KB	Sys = 8209 KB	HeapAlloc = 210 KB	HeapSys = 3808 KB	HeapIdle = 3024 KB
}

func readMemStats() {
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)

	fmt.Printf("Alloc = %v KB", bToKB(ms.Alloc))
	fmt.Printf("\tTotalAlloc = %v KB", bToKB(ms.TotalAlloc))
	fmt.Printf("\tSys = %v KB", bToKB(ms.Sys))
	fmt.Printf("\tHeapAlloc = %v KB", bToKB(ms.HeapAlloc))
	fmt.Printf("\tHeapSys = %v KB", bToKB(ms.HeapSys))
	fmt.Printf("\tHeapIdle = %v KB\n", bToKB(ms.HeapIdle))
}

func bToKB(b uint64) uint64 {
	return b / 1024 // 1024 * 1024 = 1048576
}

func TestGoroutineWithBody(t *testing.T) {
	fmt.Println("before goroutine num:", runtime.NumGoroutine())
	readMemStats()

	client := http.Client{
		Timeout: time.Second * 5,
	}
	for i := 0; i < 5; i++ {
		resp, err := client.Get("https://www.baidu.com")
		if err != nil {
			fmt.Println(err)
			continue
		}
		_, _ = io.ReadAll(resp.Body)
		//resp.Body.Close()
		//调用 resp.Body.Close()的协成数和内存
		//before goroutine num: 2
		//Alloc = 205 KB	TotalAlloc = 205 KB	Sys = 8209 KB	HeapAlloc = 205 KB	HeapSys = 3840 KB	HeapIdle = 3184 KB
		//after goroutine num: 4
		//Alloc = 622 KB	TotalAlloc = 622 KB	Sys = 8209 KB	HeapAlloc = 622 KB	HeapSys = 3808 KB	HeapIdle = 2176 KB

		//resp.Body.Close()
		//未调用 resp.Body.Close()的协成数和内存
		//before goroutine num: 2
		//Alloc = 205 KB	TotalAlloc = 205 KB	Sys = 8209 KB	HeapAlloc = 205 KB	HeapSys = 3808 KB	HeapIdle = 3024 KB
		//after goroutine num: 4
		//Alloc = 622 KB	TotalAlloc = 622 KB	Sys = 8209 KB	HeapAlloc = 622 KB	HeapSys = 3744 KB	HeapIdle = 2064 KB
	}

	time.Sleep(time.Second * 5)
	fmt.Println("after goroutine num:", runtime.NumGoroutine())
	readMemStats()
	// Output:
	//before goroutine num: 2
	//after goroutine num: 4
}
