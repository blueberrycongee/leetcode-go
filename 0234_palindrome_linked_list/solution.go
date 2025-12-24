package palindromelinkedlist

// 题目：234. 回文链表
// 链接：https://leetcode.cn/problems/palindrome-linked-list/
// 难度：Easy

// 思路：快慢指针找中点 + 反转后半部分 + 比较
// 1. 快慢指针找到链表中点
// 2. 反转后半部分链表
// 3. 比较前后两半

// 时间复杂度：O(n)
// 空间复杂度：O(1)

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// 找中点
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 反转后半部分
	secondHalf := reverseList(slow.Next)

	// 比较
	p1, p2 := head, secondHalf
	for p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	return true
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
