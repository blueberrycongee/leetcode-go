package intersectionoftwolinkedlists

import "testing"

func TestGetIntersectionNode(t *testing.T) {
	// 测试1：有相交
	t.Run("有相交", func(t *testing.T) {
		// 公共部分 [8,4,5]
		common := &ListNode{Val: 8, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}
		// A: [4,1] -> common
		headA := &ListNode{Val: 4, Next: &ListNode{Val: 1, Next: common}}
		// B: [5,6,1] -> common
		headB := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 1, Next: common}}}

		result := getIntersectionNode(headA, headB)
		if result != common {
			t.Errorf("expected intersection at node with val 8")
		}
	})

	// 测试2：无相交
	t.Run("无相交", func(t *testing.T) {
		headA := &ListNode{Val: 1, Next: &ListNode{Val: 2}}
		headB := &ListNode{Val: 3, Next: &ListNode{Val: 4}}

		result := getIntersectionNode(headA, headB)
		if result != nil {
			t.Errorf("expected nil, got %v", result)
		}
	})

	// 测试3：空链表
	t.Run("空链表", func(t *testing.T) {
		result := getIntersectionNode(nil, nil)
		if result != nil {
			t.Errorf("expected nil")
		}
	})
}
