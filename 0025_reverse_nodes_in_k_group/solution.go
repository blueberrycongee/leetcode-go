package reversenodesinkgroup

// 题目：25. K个一组翻转链表
// 链接：https://leetcode.cn/problems/reverse-nodes-in-k-group/
// 难度：Hard

// 思路：分组翻转
// 1. 检查剩余节点是否够 k 个
// 2. 翻转当前 k 个节点
// 3. 递归处理剩余部分
// 4. 连接翻转后的组

// 时间复杂度：O(n)
// 空间复杂度：O(n/k) 递归栈

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	// 检查是否有 k 个节点
	curr := head
	count := 0
	for curr != nil && count < k {
		curr = curr.Next
		count++
	}

	if count < k {
		return head // 不足 k 个，不翻转
	}

	// 翻转前 k 个
	newHead := reverse(head, k)

	// 递归处理剩余部分，head 现在是翻转后的尾部
	head.Next = reverseKGroup(curr, k)

	return newHead
}

// 翻转前 k 个节点，返回新头
func reverse(head *ListNode, k int) *ListNode {
	var prev *ListNode
	curr := head
	for i := 0; i < k; i++ {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
