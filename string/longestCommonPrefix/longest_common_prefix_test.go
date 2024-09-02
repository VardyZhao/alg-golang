package longestCommonPrefix

import (
	"fmt"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {

	strs := []string{"dog", "racecar", "car"}
	fmt.Println(LongestCommonPrefix(strs))
}
