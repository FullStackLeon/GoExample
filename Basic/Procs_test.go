package Basic

import (
	"fmt"
	"runtime"
	"testing"
)

func TestProcs(t *testing.T) {
	fmt.Println("cpu num:", runtime.NumCPU())
	fmt.Println("default procs num:", runtime.GOMAXPROCS(0))
	newValue := 100000
	previousProcs := runtime.GOMAXPROCS(newValue)
	fmt.Println("previousProcs:", previousProcs)
	fmt.Println("currentProcs after setting:", runtime.GOMAXPROCS(0))
}
