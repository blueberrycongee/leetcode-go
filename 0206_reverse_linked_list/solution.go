package reverselinkedlist

// 题目：206. 反转链表
// 链接：https://leetcode.cn/problems/reverse-linked-list/
// 难度：Easy

// 思路：迭代法
// 1. 用 prev 保存前一个节点
// 2. 遍历链表，每个节点的 next 指向 prev
// 3. 移动 prev 和 curr

// 时间复杂度：O(n)
// 空间复杂度：O(1)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		next := curr.Next // 保存下一个
		curr.Next = prev  // 反转
		prev = curr       // 移动 prev
		curr = next       // 移动 curr
	}

	return prev
}
