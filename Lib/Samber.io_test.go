package Lib

import (
	"fmt"
	"testing"

	"github.com/samber/lo"
)

func TestSamber(t *testing.T) {
	filtered := lo.Filter([]int{1, 2, 3, 4, 5}, func(item, index int) bool { return item%2 == 0 })
	fmt.Println(filtered) // [2, 4]

	doubled := lo.Map([]int{1, 2, 3, 4, 5}, func(item, index int) int {
		return item * 2
	})
	fmt.Println(doubled) // [2 4 6 8 10]

	doubledMap := lo.FilterMap([]int{1, 2, 3, 4, 5}, func(item, index int) (int, bool) {
		if item%2 == 0 {
			return item * 2, true
		}
		return item, false
	})

	fmt.Println(doubledMap) // [4 8]

	flatMap := lo.FlatMap([]int{1, 2, 3, 4, 5}, func(item int, index int) []int {
		return []int{item * 2}
	})
	fmt.Println(flatMap)
}
