package merge

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {

	nums1 := []int{0}
	m := 0
	nums2 := []int{1}
	n := 1
	Merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}
