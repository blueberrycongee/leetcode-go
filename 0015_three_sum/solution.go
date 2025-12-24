package threesum

import "slices"

// 题目：15. 三数之和
// 链接：https://leetcode.cn/problems/3sum/
// 难度：Medium
// 描述：给你一个整数数组 nums，判断是否存在三元组 [nums[i], nums[j], nums[k]]
// 满足 i != j、i != k 且 j != k，同时还满足 nums[i] + nums[j] + nums[k] == 0。
// 返回所有和为 0 且不重复的三元组。
// 注意：答案中不可以包含重复的三元组。

func threeSum(nums []int) [][]int {
	result := [][]int{}
	n := len(nums)
	if n < 3 {
		return result
	}

	// 1. 先排序，方便去重 + 双指针
	slices.Sort(nums)

	// 2. 固定第一个数 nums[i]，剩下两个数用双指针找
	for i := 0; i < n-2; i++ {
		// 剪枝：如果当前数 > 0，后面都是正数，不可能和为0
		if nums[i] > 0 {
			break
		}

		// 去重：跳过重复的第一个数
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// 双指针找剩下两个数
		left, right := i+1, n-1
		target := -nums[i] // 需要找的两数之和

		for left < right {
			sum := nums[left] + nums[right]

			if sum == target {
				// 找到一组解
				result = append(result, []int{nums[i], nums[left], nums[right]})

				// 去重：跳过重复的 left 和 right
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				// 继续找下一组
				left++
				right--
			} else if sum < target {
				left++ // 和太小，左指针右移
			} else {
				right-- // 和太大，右指针左移
			}
		}
	}

	return result
}
