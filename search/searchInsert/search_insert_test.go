package searchInsert

import (
	"fmt"
	"testing"
)

func TestSearchInsert(t *testing.T) {
	nums := []int{1, 3, 5, 6, 8, 10, 12, 14}
	target := 15
	fmt.Println(SearchInsert(nums, target))
}
