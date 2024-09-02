package inorderTraversal

/*
给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。



示例 1：
输入：root = [1,null,2,3]
输出：[1,3,2]

示例 2：
输入：root = []
输出：[]

示例 3：
输入：root = [1]
输出：[1]


提示：

树中节点数目在范围 [0, 100] 内
-100 <= Node.val <= 100


进阶: 递归算法很简单，你可以通过迭代算法完成吗？
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InorderTraversal(root *TreeNode) []int {
	// 按层级逐层输出，每一层输出顺序为左中右
	res := make([]int, 0)
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}

	//var dfs func(root *TreeNode)
	//dfs = func(node *TreeNode) {
	//	if node == nil {
	//		return
	//	}
	//	dfs(node.Left)
	//	res = append(res, node.Val)
	//	dfs(node.Right)
	//}
	//dfs(root)

	return res
}
