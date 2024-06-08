package Lock

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type Counter struct {
	count int
	sync.Mutex
}

func TestMutexLock(t *testing.T) {
	c := Counter{}
	wg := &sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Lock()
			defer c.Unlock()
			c.count++
		}()
	}
	wg.Wait()
	fmt.Println("count: ", c.count)
}

func TestMutexTryLock(t *testing.T) {
	c := Counter{}
	wg := &sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				if c.TryLock() {
					defer c.Unlock()
					c.count++
					return
				}
				time.Sleep(time.Millisecond * 100)
			}
		}()
	}
	wg.Wait()
	fmt.Println("count: ", c.count)
}
