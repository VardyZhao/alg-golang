package streamMovingAverage

import (
	"fmt"
	"testing"
)

func TestStreamMovingAverage(t *testing.T) {

	size := 3
	values := []int{1, 10, 3, 5}
	obj := Constructor(size)
	for _, v := range values {
		fmt.Println(obj.Next(v))
	}
}
