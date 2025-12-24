package longestvalidparentheses

// 题目：32. 最长有效括号
// 链接：https://leetcode.cn/problems/longest-valid-parentheses/
// 难度：Hard

// 思路：栈
// 1. 栈中存放下标，栈底始终保持最后一个没被匹配的右括号下标
// 2. 遇到左括号，入栈
// 3. 遇到右括号，先弹出，如果栈空则当前下标入栈，否则计算长度

// 时间复杂度：O(n)
// 空间复杂度：O(n)

func longestValidParentheses(s string) int {
	maxLen := 0
	stack := []int{-1} // 初始化，作为边界

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1] // 弹出
			if len(stack) == 0 {
				stack = append(stack, i) // 当前右括号作为新的边界
			} else {
				length := i - stack[len(stack)-1]
				if length > maxLen {
					maxLen = length
				}
			}
		}
	}

	return maxLen
}
