package containerwithmostwater

// 题目：11. 盛最多水的容器
// 链接：https://leetcode.cn/problems/container-with-most-water/
// 难度：Medium
// 描述：给定一个长度为 n 的整数数组 height。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i])。
// 找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
// 返回容器可以储存的最大水量。

func maxArea(height []int) int {
	// 双指针：左右夹逼
	left, right := 0, len(height)-1
	maxWater := 0

	for left < right {
		// 宽度 = right - left
		// 高度 = 较矮的那条线（木桶效应）
		width := right - left
		h := min(height[left], height[right])
		water := width * h

		// 更新最大值
		if water > maxWater {
			maxWater = water
		}

		// 关键：移动较矮的那边
		// 因为宽度在缩小，只有找到更高的线才可能翻盘
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxWater
}
