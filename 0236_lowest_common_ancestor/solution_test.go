package lowestcommonancestor

import "testing"

func TestLowestCommonAncestor(t *testing.T) {
	// 构建测试树
	//       3
	//      / \
	//     5   1
	//    / \ / \
	//   6  2 0  8
	//     / \
	//    7   4
	node7 := &TreeNode{Val: 7}
	node4 := &TreeNode{Val: 4}
	node6 := &TreeNode{Val: 6}
	node2 := &TreeNode{Val: 2, Left: node7, Right: node4}
	node0 := &TreeNode{Val: 0}
	node8 := &TreeNode{Val: 8}
	node5 := &TreeNode{Val: 5, Left: node6, Right: node2}
	node1 := &TreeNode{Val: 1, Left: node0, Right: node8}
	root := &TreeNode{Val: 3, Left: node5, Right: node1}

	t.Run("LCA(5,1)=3", func(t *testing.T) {
		result := lowestCommonAncestor(root, node5, node1)
		if result != root {
			t.Errorf("expected root(3), got %v", result.Val)
		}
	})

	t.Run("LCA(5,4)=5", func(t *testing.T) {
		result := lowestCommonAncestor(root, node5, node4)
		if result != node5 {
			t.Errorf("expected node5, got %v", result.Val)
		}
	})

	t.Run("LCA(6,4)=5", func(t *testing.T) {
		result := lowestCommonAncestor(root, node6, node4)
		if result != node5 {
			t.Errorf("expected node5, got %v", result.Val)
		}
	})
}
