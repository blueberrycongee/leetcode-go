package uniquepaths

// 题目：62. 不同路径
// 链接：https://leetcode.cn/problems/unique-paths/
// 难度：Medium

// 思路：动态规划
// dp[i][j] = dp[i-1][j] + dp[i][j-1]
// 空间优化：只用一维数组

// 时间复杂度：O(m*n)
// 空间复杂度：O(n)

func uniquePaths(m int, n int) int {
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1 // 第一行都是1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] += dp[j-1]
		}
	}

	return dp[n-1]
}
