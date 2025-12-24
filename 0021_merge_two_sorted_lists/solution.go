package mergetwosortedlists

// 题目：21. 合并两个有序链表
// 链接：https://leetcode.cn/problems/merge-two-sorted-lists/
// 难度：Easy

// 思路：迭代 + 哨兵节点
// 1. 创建哨兵节点简化边界处理
// 2. 比较两链表当前节点，较小的接入结果链表
// 3. 将剩余链表接到尾部

// 时间复杂度：O(m+n)
// 空间复杂度：O(1)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}
		curr = curr.Next
	}

	if list1 != nil {
		curr.Next = list1
	} else {
		curr.Next = list2
	}

	return dummy.Next
}
