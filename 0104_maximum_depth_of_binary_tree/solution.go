package maximumdepthofbinarytree

// 题目：104. 二叉树的最大深度
// 链接：https://leetcode.cn/problems/maximum-depth-of-binary-tree/
// 难度：Easy

// 思路：递归
// 深度 = 1 + max(左子树深度, 右子树深度)

// 时间复杂度：O(n)
// 空间复杂度：O(h)，h为树高

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}
