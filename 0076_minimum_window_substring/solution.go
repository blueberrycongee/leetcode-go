package minimumwindowsubstring

// 题目：76. 最小覆盖子串
// 链接：https://leetcode.cn/problems/minimum-window-substring/
// 难度：Hard

// 思路：滑动窗口
// 1. 用 need 记录 t 中每个字符的需求量
// 2. 用 window 记录当前窗口中的字符计数
// 3. 右指针扩展直到满足条件，左指针收缩寻找最小窗口

// 时间复杂度：O(n)
// 空间复杂度：O(字符集大小)

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	need := make(map[byte]int)
	window := make(map[byte]int)

	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	left, right := 0, 0
	valid := 0 // 已满足的字符种类数
	start, minLen := 0, len(s)+1

	for right < len(s) {
		c := s[right]
		right++

		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		// 当所有字符都满足时，收缩左边界
		for valid == len(need) {
			if right-left < minLen {
				start = left
				minLen = right - left
			}
			d := s[left]
			left++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}

	if minLen == len(s)+1 {
		return ""
	}
	return s[start : start+minLen]
}
