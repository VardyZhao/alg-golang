package maxDepth

/*
给定一个二叉树 root ，返回其最大深度。

二叉树的 最大深度 是指从根节点到最远叶子节点的最长路径上的节点数。



示例 1：
		3
	9		20
		15		7
输入：root = [3,9,20,null,null,15,7]
输出：3

示例 2：
		1
			2
输入：root = [1,null,2]
输出：2


提示：

树中节点的数量在 [0, 104] 区间内。
-100 <= Node.val <= 100
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MaxDepth(root *TreeNode) int {

	// todo

	depth := 0
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		depth = max(depth, len(stack))
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = root.Right
	}

	return depth
}
