package trappingrainwater

// 题目：42. 接雨水
// 链接：https://leetcode.cn/problems/trapping-rain-water/
// 难度：Hard

// 思路：双指针
// 1. 左右指针从两端向中间移动
// 2. 维护左边最大高度 leftMax 和右边最大高度 rightMax
// 3. 每个位置能接的水 = min(leftMax, rightMax) - height[i]
// 4. 哪边较矮就移动哪边（因为较矮的一边决定了水位）

// 时间复杂度：O(n)
// 空间复杂度：O(1)

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}

	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	result := 0

	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				result += leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				result += rightMax - height[right]
			}
			right--
		}
	}

	return result
}
