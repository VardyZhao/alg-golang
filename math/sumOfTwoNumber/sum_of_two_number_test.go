package sumOfTwoNumber

import (
	"fmt"
	"testing"
)

func TestSumOfTwoNumber(t *testing.T) {

	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(SumOfTwoNumber(nums, target))
}
