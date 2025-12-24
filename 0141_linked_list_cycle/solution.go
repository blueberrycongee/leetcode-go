package linkedlistcycle

// 题目：141. 环形链表
// 链接：https://leetcode.cn/problems/linked-list-cycle/
// 难度：Easy

// 思路：快慢指针（Floyd判圈）
// 1. 慢指针每次走一步，快指针每次走两步
// 2. 如果有环，快慢指针必定相遇
// 3. 如果无环，快指针会先到达 nil

// 时间复杂度：O(n)
// 空间复杂度：O(1)

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}

	return false
}
