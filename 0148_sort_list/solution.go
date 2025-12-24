package sortlist

// 题目：148. 排序链表
// 链接：https://leetcode.cn/problems/sort-list/
// 难度：Medium

// 思路：归并排序
// 1. 快慢指针找中点，断开链表
// 2. 递归排序左右两半
// 3. 合并两个有序链表

// 时间复杂度：O(nlogn)
// 空间复杂度：O(logn) 递归栈

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 找中点并断开
	mid := getMid(head)
	left := head
	right := mid.Next
	mid.Next = nil

	// 递归排序
	left = sortList(left)
	right = sortList(right)

	// 合并
	return merge(left, right)
}

func getMid(head *ListNode) *ListNode {
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func merge(l1, l2 *ListNode) *ListNode {
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
