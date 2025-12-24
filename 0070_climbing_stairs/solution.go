package climbingstairs

// 题目：70. 爬楼梯
// 链接：https://leetcode.cn/problems/climbing-stairs/
// 难度：Easy

// 思路：动态规划（斐波那契）
// dp[i] = dp[i-1] + dp[i-2]
// 到达第 i 阶 = 从 i-1 阶走一步 + 从 i-2 阶走两步

// 时间复杂度：O(n)
// 空间复杂度：O(1)

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	prev, curr := 1, 2
	for i := 3; i <= n; i++ {
		prev, curr = curr, prev+curr
	}
	return curr
}
