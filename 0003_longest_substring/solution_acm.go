package main

import (
	"bufio"
	"fmt"
	"os"
)

// 题目：3. 无重复字符的最长子串 (ACM模式)
// 输入格式：
// 一行字符串 s
// 输出格式：
// 一个整数，表示无重复字符的最长子串的长度

// 思路：滑动窗口 + 哈希表
// 1. 用 map 记录每个字符最后出现的位置
// 2. 右指针遍历，左指针在遇到重复字符时跳到重复位置+1
// 3. 每步更新最大长度

// 时间复杂度：O(n)
// 空间复杂度：O(min(m,n))，m为字符集大小

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	
	// 读取输入字符串
	scanner.Scan()
	s := scanner.Text()
	
	// 计算结果
	result := lengthOfLongestSubstring(s)
	
	// 输出结果
	fmt.Println(result)
}

func lengthOfLongestSubstring(s string) int {
	charIndex := make(map[byte]int)
	maxLen := 0
	left := 0

	for right := 0; right < len(s); right++ {
		ch := s[right]
		if idx, ok := charIndex[ch]; ok && idx >= left {
			left = idx + 1
		}
		charIndex[ch] = right
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}

	return maxLen
}
