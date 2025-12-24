package linkedlistcycleii

import "testing"

func TestDetectCycle(t *testing.T) {
	// 测试1：有环
	t.Run("有环-入口在第二个节点", func(t *testing.T) {
		node1 := &ListNode{Val: 3}
		node2 := &ListNode{Val: 2}
		node3 := &ListNode{Val: 0}
		node4 := &ListNode{Val: -4}
		node1.Next = node2
		node2.Next = node3
		node3.Next = node4
		node4.Next = node2

		result := detectCycle(node1)
		if result != node2 {
			t.Error("expected node2")
		}
	})

	// 测试2：两节点有环
	t.Run("两节点有环", func(t *testing.T) {
		node1 := &ListNode{Val: 1}
		node2 := &ListNode{Val: 2}
		node1.Next = node2
		node2.Next = node1

		result := detectCycle(node1)
		if result != node1 {
			t.Error("expected node1")
		}
	})

	// 测试3：无环
	t.Run("无环", func(t *testing.T) {
		node1 := &ListNode{Val: 1}

		result := detectCycle(node1)
		if result != nil {
			t.Error("expected nil")
		}
	})
}
