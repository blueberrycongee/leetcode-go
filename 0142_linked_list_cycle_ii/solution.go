package linkedlistcycleii

// 题目：142. 环形链表 II
// 链接：https://leetcode.cn/problems/linked-list-cycle-ii/
// 难度：Medium

// 思路：快慢指针 + 数学推导
// 1. 快慢指针找到相遇点
// 2. 相遇后，一个指针从头开始，另一个从相遇点开始
// 3. 两指针每次走一步，相遇点即为环的入口
// 数学证明：设头到入口距离a，入口到相遇点b，环长c
// 快指针走 a+b+nc，慢指针走 a+b，快=2慢，得 a = nc - b = (n-1)c + (c-b)

// 时间复杂度：O(n)
// 空间复杂度：O(1)

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow, fast := head, head

	// 找相遇点
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			// 找入口
			ptr := head
			for ptr != slow {
				ptr = ptr.Next
				slow = slow.Next
			}
			return ptr
		}
	}

	return nil
}
