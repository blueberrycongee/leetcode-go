package intersectionoftwolinkedlists

// 题目：160. 相交链表
// 链接：https://leetcode.cn/problems/intersection-of-two-linked-lists/
// 难度：Easy

// 思路：双指针
// 1. 指针 A 遍历完链表 A 后，转到链表 B 的头部继续遍历
// 2. 指针 B 遍历完链表 B 后，转到链表 A 的头部继续遍历
// 3. 如果相交，两指针会在交点相遇（走过的路程相同）
// 4. 如果不相交，两指针会同时到达 nil

// 时间复杂度：O(m+n)
// 空间复杂度：O(1)

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	pA, pB := headA, headB

	for pA != pB {
		if pA == nil {
			pA = headB
		} else {
			pA = pA.Next
		}
		if pB == nil {
			pB = headA
		} else {
			pB = pB.Next
		}
	}

	return pA
}
