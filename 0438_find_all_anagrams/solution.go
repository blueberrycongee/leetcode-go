package findallanagrams

// 题目：438. 找到字符串中所有字母异位词
// 链接：https://leetcode.cn/problems/find-all-anagrams-in-a-string/
// 难度：Medium

// 思路：滑动窗口 + 字符计数
// 1. 统计 p 中每个字符的出现次数
// 2. 用固定大小的滑动窗口遍历 s
// 3. 当窗口内字符计数与 p 相同时，记录起始位置

// 时间复杂度：O(n)
// 空间复杂度：O(1)，只用26个字母

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return nil
	}

	var result []int
	var sCount, pCount [26]int

	// 统计 p 的字符频率
	for i := 0; i < len(p); i++ {
		pCount[p[i]-'a']++
		sCount[s[i]-'a']++
	}

	if sCount == pCount {
		result = append(result, 0)
	}

	// 滑动窗口
	for i := len(p); i < len(s); i++ {
		sCount[s[i]-'a']++        // 加入右边
		sCount[s[i-len(p)]-'a']-- // 移除左边
		if sCount == pCount {
			result = append(result, i-len(p)+1)
		}
	}

	return result
}
