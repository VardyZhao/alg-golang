package inorderTraversal

import (
	"fmt"
	"testing"
)

func TestInorderTraversal(t *testing.T) {

	list := []interface{}{1, nil, 2, 3}
	//list := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	tree := buildTree(list)
	fmt.Println(InorderTraversal(tree))
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
