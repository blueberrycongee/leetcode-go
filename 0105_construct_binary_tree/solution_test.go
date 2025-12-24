package constructbinarytree

import (
	"reflect"
	"testing"
)

// 层序遍历验证
func levelOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		result = append(result, node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return result
}

func TestBuildTree(t *testing.T) {
	tests := []struct {
		name     string
		preorder []int
		inorder  []int
		expected []int // 层序遍历结果
	}{
		{
			name:     "示例1",
			preorder: []int{3, 9, 20, 15, 7},
			inorder:  []int{9, 3, 15, 20, 7},
			expected: []int{3, 9, 20, 15, 7},
		},
		{
			name:     "单节点",
			preorder: []int{-1},
			inorder:  []int{-1},
			expected: []int{-1},
		},
		{
			name:     "空树",
			preorder: []int{},
			inorder:  []int{},
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTree(tt.preorder, tt.inorder)
			result := levelOrder(root)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
