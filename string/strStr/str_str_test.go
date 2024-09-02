package strStr

import (
	"fmt"
	"testing"
)

func TestStrStr(t *testing.T) {
	haystack, needle := "sadbutsad", "adb"
	fmt.Println(StrStr(haystack, needle))
}
