package lengthOfLongestSubstring

import (
	"fmt"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	s := "abcadcbb"
	res := LengthOfLongestSubstring(s)
	fmt.Println(res)
}
