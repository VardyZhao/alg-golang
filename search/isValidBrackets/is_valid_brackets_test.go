package isValidBrackets

import (
	"fmt"
	"testing"
)

func TestIsValidBrackets(t *testing.T) {

	s := "{{[[[]()]]}}"
	fmt.Println(IsValidBrackets(s))
}
