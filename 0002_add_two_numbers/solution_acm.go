package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 题目：2. 两数相加 (ACM模式)
// 输入格式：
// 第一行：用空格分隔的整数，表示第一个数字的各位（逆序）
// 第二行：用空格分隔的整数，表示第二个数字的各位（逆序）
// 输出格式：
// 一行包含用空格分隔的整数，表示结果数字的各位（逆序）

// 思路：模拟加法
// 1. 同时遍历两个数字序列，逐位相加
// 2. 维护进位 carry
// 3. 注意最后可能还有进位

// 时间复杂度：O(max(m,n))
// 空间复杂度：O(max(m,n))

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	
	// 读取第一行：第一个数字的各位（逆序）
	scanner.Scan()
	nums1Str := strings.Fields(scanner.Text())
	nums1 := make([]int, len(nums1Str))
	for i, s := range nums1Str {
		nums1[i], _ = strconv.Atoi(s)
	}
	
	// 读取第二行：第二个数字的各位（逆序）
	scanner.Scan()
	nums2Str := strings.Fields(scanner.Text())
	nums2 := make([]int, len(nums2Str))
	for i, s := range nums2Str {
		nums2[i], _ = strconv.Atoi(s)
	}
	
	// 构建链表
	l1 := buildList(nums1)
	l2 := buildList(nums2)
	
	// 执行两数相加
	result := addTwoNumbers(l1, l2)
	
	// 输出结果
	printList(result)
}

func buildList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	
	head := &ListNode{Val: nums[0]}
	curr := head
	for i := 1; i < len(nums); i++ {
		curr.Next = &ListNode{Val: nums[i]}
		curr = curr.Next
	}
	return head
}

func printList(head *ListNode) {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	
	for i, val := range result {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(val)
	}
	fmt.Println()
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