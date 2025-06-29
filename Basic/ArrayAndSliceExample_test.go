package Basic

import (
	"fmt"
	"testing"
)

func TestArrayToSlice(t *testing.T) {
	arr := [3]int{1, 2, 3}
	fmt.Printf("type of arr: %T\n", arr)

	slice := arr[:]
	fmt.Printf("type of slice: %T\n", slice)
}

// 切片转数组时，切片的容量要大于等于数组的长度，否则会panic
func TestSliceToPointer(t *testing.T) {
	slice := []int{1, 2, 3}
	fmt.Printf("type of slice: %T\n", slice)

	arrPtr := (*[3]int)(slice)
	fmt.Printf("type of arrPtr: %T\n", arrPtr)
	arrPtr[0] = 100
	fmt.Printf("arr: %v\n", *arrPtr)
	fmt.Printf("slice: %v\n", slice)
}
