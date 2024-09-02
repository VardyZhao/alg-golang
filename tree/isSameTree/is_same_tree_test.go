package isSameTree

import (
	"fmt"
	"testing"
)

func TestIsSameTree(t *testing.T) {

	p := []interface{}{1, 3}
	q := []interface{}{1, nil, 3}
	pTree := buildTree(p)
	qTree := buildTree(q)
	fmt.Println(IsSameTree(pTree, qTree))
}

func buildTree(list []interface{}) *TreeNode {
	if len(list) == 0 {
		return nil
	}
	root := &TreeNode{Val: list[0].(int)}
	queue := []*TreeNode{root}
	i := 1
	for i < len(list) {
		cur := queue[0]
		queue = queue[1:]
		if i < len(list) && list[i] != nil {
			cur.Left = &TreeNode{Val: list[i].(int)}
			queue = append(queue, cur.Left)
		}
		i++

		if i < len(list) && list[i] != nil {
			cur.Right = &TreeNode{Val: list[i].(int)}
			queue = append(queue, cur.Right)
		}
		i++
	}
	return root
}
