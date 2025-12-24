package constructbinarytree

// 题目：105. 从前序与中序遍历序列构造二叉树
// 链接：https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
// 难度：Medium

// 思路：递归
// 1. 前序第一个元素是根节点
// 2. 在中序中找到根节点位置，划分左右子树
// 3. 递归构建左右子树

// 时间复杂度：O(n)
// 空间复杂度：O(n)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	// 用哈希表加速查找
	inorderIdx := make(map[int]int)
	for i, v := range inorder {
		inorderIdx[v] = i
	}

	return build(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1, inorderIdx)
}

func build(preorder []int, preStart, preEnd int, inorder []int, inStart, inEnd int, inorderIdx map[int]int) *TreeNode {
	if preStart > preEnd {
		return nil
	}

	rootVal := preorder[preStart]
	root := &TreeNode{Val: rootVal}

	rootIdx := inorderIdx[rootVal]
	leftSize := rootIdx - inStart

	root.Left = build(preorder, preStart+1, preStart+leftSize, inorder, inStart, rootIdx-1, inorderIdx)
	root.Right = build(preorder, preStart+leftSize+1, preEnd, inorder, rootIdx+1, inEnd, inorderIdx)

	return root
}
