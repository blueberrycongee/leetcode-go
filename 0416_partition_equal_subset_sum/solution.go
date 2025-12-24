package partitionequalsubsetsum

// 题目：416. 分割等和子集
// 链接：https://leetcode.cn/problems/partition-equal-subset-sum/
// 难度：Medium

// 思路：0-1背包
// 问题转化为：能否从数组中选出一些数，使其和等于 sum/2
// dp[j] = 能否凑出和为 j

// 时间复杂度：O(n * sum)
// 空间复杂度：O(sum)

func canPartition(nums []int) bool {
	sum := 0
	for _, n := range nums {
		sum += n
	}

	if sum%2 != 0 {
		return false
	}

	target := sum / 2
	dp := make([]bool, target+1)
	dp[0] = true

	for _, num := range nums {
		// 倒序遍历，避免重复使用同一个数
		for j := target; j >= num; j-- {
			dp[j] = dp[j] || dp[j-num]
		}
	}

	return dp[target]
}
