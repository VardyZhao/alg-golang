package deleteDuplicates

import (
	"fmt"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {

	array := []int{1, 1, 2, 3, 3}
	var list *ListNode
	var nextNode *ListNode
	for i := len(array) - 1; i >= 0; i-- {
		list = &ListNode{Val: array[i], Next: nextNode}
		nextNode = list
	}

	fmt.Println(DeleteDuplicates(list))
}
