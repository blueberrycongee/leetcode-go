package wordbreak

// 题目：139. 单词拆分
// 链接：https://leetcode.cn/problems/word-break/
// 难度：Medium

// 思路：动态规划
// dp[i] = s[0:i] 能否被拆分
// dp[i] = dp[j] && s[j:i] in wordDict (for all j < i)

// 时间复杂度：O(n^2)
// 空间复杂度：O(n)

func wordBreak(s string, wordDict []string) bool {
	wordSet := make(map[string]bool)
	for _, w := range wordDict {
		wordSet[w] = true
	}

	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true

	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[n]
}
