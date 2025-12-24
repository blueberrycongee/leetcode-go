package binarytreelevelordertraversal

// 题目：102. 二叉树的层序遍历
// 链接：https://leetcode.cn/problems/binary-tree-level-order-traversal/
// 难度：Medium

// 思路：BFS + 队列
// 1. 每次处理一层的所有节点
// 2. 将下一层节点加入队列

// 时间复杂度：O(n)
// 空间复杂度：O(n)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var result [][]int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		level := make([]int, 0, levelSize)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, level)
	}

	return result
}
