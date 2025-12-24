package symmetrictree

// 题目：101. 对称二叉树
// 链接：https://leetcode.cn/problems/symmetric-tree/
// 难度：Easy

// 思路：递归比较镜像
// 1. 左子树的左 == 右子树的右
// 2. 左子树的右 == 右子树的左

// 时间复杂度：O(n)
// 空间复杂度：O(h)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

func isMirror(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	return left.Val == right.Val &&
		isMirror(left.Left, right.Right) &&
		isMirror(left.Right, right.Left)
}
