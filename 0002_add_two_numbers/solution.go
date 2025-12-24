package addtwonumbers

// 题目：2. 两数相加
// 链接：https://leetcode.cn/problems/add-two-numbers/
// 难度：Medium

// 思路：模拟加法
// 1. 同时遍历两个链表，逐位相加
// 2. 维护进位 carry
// 3. 注意最后可能还有进位

// 时间复杂度：O(max(m,n))
// 空间复杂度：O(max(m,n))

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10
		curr.Next = &ListNode{Val: sum % 10}
		curr = curr.Next
	}

	return dummy.Next
}
