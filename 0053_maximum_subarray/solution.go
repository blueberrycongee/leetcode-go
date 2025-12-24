package maximumsubarray

// 题目：53. 最大子数组和
// 链接：https://leetcode.cn/problems/maximum-subarray/
// 难度：Medium

// 思路：动态规划（Kadane算法）
// dp[i] = max(nums[i], dp[i-1] + nums[i])
// 要么从当前开始，要么接上之前的

// 时间复杂度：O(n)
// 空间复杂度：O(1)

func maxSubArray(nums []int) int {
	maxSum := nums[0]
	currSum := nums[0]

	for i := 1; i < len(nums); i++ {
		if currSum < 0 {
			currSum = nums[i]
		} else {
			currSum += nums[i]
		}
		if currSum > maxSum {
			maxSum = currSum
		}
	}

	return maxSum
}
