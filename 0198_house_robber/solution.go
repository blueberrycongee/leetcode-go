package houserobber

// 题目：198. 打家劫舍
// 链接：https://leetcode.cn/problems/house-robber/
// 难度：Medium

// 思路：动态规划
// dp[i] = max(dp[i-1], dp[i-2] + nums[i])
// 要么不偷当前房，要么偷当前房（必须跳过前一个）

// 时间复杂度：O(n)
// 空间复杂度：O(1)

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	prev, curr := nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		prev, curr = curr, max(curr, prev+nums[i])
	}
	return curr
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
