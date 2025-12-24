package invertbinarytree

// 题目：226. 翻转二叉树
// 链接：https://leetcode.cn/problems/invert-binary-tree/
// 难度：Easy

// 思路：递归
// 1. 交换左右子树
// 2. 递归翻转左右子树

// 时间复杂度：O(n)
// 空间复杂度：O(h)，h为树高

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}
