package coinchange

// 题目：322. 零钱兑换
// 链接：https://leetcode.cn/problems/coin-change/
// 难度：Medium

// 思路：动态规划（完全背包）
// dp[i] = 凑出金额 i 所需的最少硬币数
// dp[i] = min(dp[i], dp[i-coin] + 1)

// 时间复杂度：O(amount * len(coins))
// 空间复杂度：O(amount)

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1 // 初始化为不可能的大值
	}

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin <= i && dp[i-coin]+1 < dp[i] {
				dp[i] = dp[i-coin] + 1
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}
