package mergeTwoLists

import (
	"fmt"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {

	array1 := []int{1, 2, 4}
	array2 := []int{1, 3, 4}

	var list1 *ListNode
	var nextNode1 *ListNode
	for i := len(array1) - 1; i >= 0; i-- {
		list1 = &ListNode{Val: array1[i], Next: nextNode1}
		nextNode1 = list1
	}

	var list2 *ListNode
	var nextNode2 *ListNode
	for i := len(array2) - 1; i >= 0; i-- {
		list2 = &ListNode{Val: array2[i], Next: nextNode2}
		nextNode2 = list2
	}

	fmt.Println(MergeTwoLists(list1, list2))
}
