package maximumproductsubarray

// 题目：152. 乘积最大子数组
// 链接：https://leetcode.cn/problems/maximum-product-subarray/
// 难度：Medium

// 思路：动态规划
// 因为负数乘负数会变正，所以需要同时维护最大值和最小值
// maxProd = max(num, maxProd*num, minProd*num)
// minProd = min(num, maxProd*num, minProd*num)

// 时间复杂度：O(n)
// 空间复杂度：O(1)

func maxProduct(nums []int) int {
	maxProd, minProd, result := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		num := nums[i]
		// 注意：要先保存旧的 maxProd
		tmpMax := max(num, max(maxProd*num, minProd*num))
		minProd = min(num, min(maxProd*num, minProd*num))
		maxProd = tmpMax

		if maxProd > result {
			result = maxProd
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
