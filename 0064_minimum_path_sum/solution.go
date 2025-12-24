package minimumpathsum

// 题目：64. 最小路径和
// 链接：https://leetcode.cn/problems/minimum-path-sum/
// 难度：Medium

// 思路：动态规划
// dp[i][j] = grid[i][j] + min(dp[i-1][j], dp[i][j-1])
// 空间优化：原地修改或一维数组

// 时间复杂度：O(m*n)
// 空间复杂度：O(1) 原地修改

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	// 初始化第一行
	for j := 1; j < n; j++ {
		grid[0][j] += grid[0][j-1]
	}

	// 初始化第一列
	for i := 1; i < m; i++ {
		grid[i][0] += grid[i-1][0]
	}

	// 填充其余
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] += min(grid[i-1][j], grid[i][j-1])
		}
	}

	return grid[m-1][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
