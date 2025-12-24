package longestincreasingsubsequence

// 题目：300. 最长递增子序列
// 链接：https://leetcode.cn/problems/longest-increasing-subsequence/
// 难度：Medium

// 思路：二分查找 + 贪心
// 维护一个数组 tails，tails[i] 表示长度为 i+1 的递增子序列的最小末尾元素
// 遍历 nums，用二分查找找到 num 应该替换的位置

// 时间复杂度：O(nlogn)
// 空间复杂度：O(n)

func lengthOfLIS(nums []int) int {
	tails := make([]int, 0, len(nums))

	for _, num := range nums {
		// 二分查找第一个 >= num 的位置
		left, right := 0, len(tails)
		for left < right {
			mid := left + (right-left)/2
			if tails[mid] < num {
				left = mid + 1
			} else {
				right = mid
			}
		}

		if left == len(tails) {
			tails = append(tails, num)
		} else {
			tails[left] = num
		}
	}

	return len(tails)
}
