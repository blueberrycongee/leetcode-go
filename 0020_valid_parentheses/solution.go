package validparentheses

// 题目：20. 有效的括号
// 链接：https://leetcode.cn/problems/valid-parentheses/
// 难度：Easy

// 思路：栈
// 遇到左括号入栈，遇到右括号检查栈顶是否匹配

// 时间复杂度：O(n)
// 空间复杂度：O(n)

func isValid(s string) bool {
	stack := make([]byte, 0)
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch == '(' || ch == '[' || ch == '{' {
			stack = append(stack, ch)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[ch] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
