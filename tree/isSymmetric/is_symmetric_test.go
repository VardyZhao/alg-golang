package isSymmetric

import (
	"fmt"
	"testing"
)

func TestIsSymmetric(t *testing.T) {

	points := []interface{}{1, 2, 2, 3, 4, 4, 3}
	//points := []interface{}{1, 2, 2, nil, 3, nil, 3}

	tree := buildTree(points)
	fmt.Println(IsSymmetric(tree))
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
