package dynamicPassword

import (
	"fmt"
	"testing"
)

func TestDynamicPassword(t *testing.T) {

	password := "s3cur1tyC0d3"
	target := 4
	fmt.Println(DynamicPassword(password, target))
}
