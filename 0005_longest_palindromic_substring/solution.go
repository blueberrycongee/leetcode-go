package longestpalindromicsubstring

// 题目：5. 最长回文子串
// 链接：https://leetcode.cn/problems/longest-palindromic-substring/
// 难度：Medium

// 思路：中心扩展
// 从每个位置向两边扩展，找最长回文
// 需要考虑奇数和偶数长度的回文

// 时间复杂度：O(n^2)
// 空间复杂度：O(1)

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	start, maxLen := 0, 1

	for i := 0; i < len(s); i++ {
		// 奇数长度
		l1 := expand(s, i, i)
		// 偶数长度
		l2 := expand(s, i, i+1)

		l := l1
		if l2 > l1 {
			l = l2
		}

		if l > maxLen {
			maxLen = l
			start = i - (l-1)/2
		}
	}

	return s[start : start+maxLen]
}

func expand(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}
