package mergeksortedlists

// 题目：23. 合并K个升序链表
// 链接：https://leetcode.cn/problems/merge-k-sorted-lists/
// 难度：Hard

// 思路：分治法
// 1. 将 k 个链表两两合并
// 2. 递归直到只剩一个链表

// 时间复杂度：O(nlogk)，n为总节点数，k为链表数
// 空间复杂度：O(logk) 递归栈

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return mergeRange(lists, 0, len(lists)-1)
}

func mergeRange(lists []*ListNode, left, right int) *ListNode {
	if left == right {
		return lists[left]
	}
	mid := left + (right-left)/2
	l1 := mergeRange(lists, left, mid)
	l2 := mergeRange(lists, mid+1, right)
	return mergeTwoLists(l1, l2)
}

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			curr.Next = l1
			l1 = l1.Next
		} else {
			curr.Next = l2
			l2 = l2.Next
		}
		curr = curr.Next
	}

	if l1 != nil {
		curr.Next = l1
	} else {
		curr.Next = l2
	}

	return dummy.Next
}
