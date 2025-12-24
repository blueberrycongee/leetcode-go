package binarytreeinordertraversal

// 题目：94. 二叉树的中序遍历
// 链接：https://leetcode.cn/problems/binary-tree-inorder-traversal/
// 难度：Easy

// 思路：迭代法（用栈模拟）
// 1. 一直向左走，将节点入栈
// 2. 弹出节点，访问它
// 3. 转向右子树

// 时间复杂度：O(n)
// 空间复杂度：O(n)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var result []int
	var stack []*TreeNode
	curr := root

	for curr != nil || len(stack) > 0 {
		// 一直向左走
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		// 弹出并访问
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, curr.Val)
		// 转向右子树
		curr = curr.Right
	}

	return result
}
