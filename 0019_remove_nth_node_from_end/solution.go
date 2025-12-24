package removenthnodefromend

// 题目：19. 删除链表的倒数第N个结点
// 链接：https://leetcode.cn/problems/remove-nth-node-from-end-of-list/
// 难度：Medium

// 思路：快慢指针
// 1. 快指针先走 n 步
// 2. 然后快慢指针同时走，直到快指针到达末尾
// 3. 此时慢指针指向倒数第 n+1 个节点
// 4. 删除慢指针的下一个节点

// 时间复杂度：O(n)
// 空间复杂度：O(1)

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	fast, slow := dummy, dummy

	// 快指针先走 n+1 步
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}

	// 同时移动
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	// 删除节点
	slow.Next = slow.Next.Next

	return dummy.Next
}
