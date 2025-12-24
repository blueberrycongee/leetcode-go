package validatebinarysearchtree

// 题目：98. 验证二叉搜索树
// 链接：https://leetcode.cn/problems/validate-binary-search-tree/
// 难度：Medium

// 思路：递归 + 上下界
// 每个节点的值必须在 (min, max) 范围内
// 左子树上界是当前节点值，右子树下界是当前节点值

// 时间复杂度：O(n)
// 空间复杂度：O(h)

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return validate(root, math.MinInt64, math.MaxInt64)
}

func validate(node *TreeNode, min, max int64) bool {
	if node == nil {
		return true
	}
	val := int64(node.Val)
	if val <= min || val >= max {
		return false
	}
	return validate(node.Left, min, val) && validate(node.Right, val, max)
}
