package lowestcommonancestor

// 题目：236. 二叉树的最近公共祖先
// 链接：https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/
// 难度：Medium

// 思路：递归后序遍历
// 1. 如果当前节点是 p 或 q，返回当前节点
// 2. 递归查找左右子树
// 3. 如果左右都找到，当前节点就是 LCA
// 4. 如果只有一边找到，返回那一边的结果

// 时间复杂度：O(n)
// 空间复杂度：O(h)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root // p 和 q 分别在左右子树
	}
	if left != nil {
		return left
	}
	return right
}
