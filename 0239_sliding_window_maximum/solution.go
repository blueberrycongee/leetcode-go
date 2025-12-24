package slidingwindowmaximum

// 题目：239. 滑动窗口最大值
// 链接：https://leetcode.cn/problems/sliding-window-maximum/
// 难度：Hard

// 思路：单调递减队列
// 1. 维护一个单调递减的双端队列，存储下标
// 2. 队首始终是当前窗口的最大值下标
// 3. 新元素入队前，移除所有比它小的元素（因为它们不可能成为最大值）
// 4. 移除过期的队首元素（下标超出窗口范围）

// 时间复杂度：O(n)
// 空间复杂度：O(k)

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return nil
	}

	result := make([]int, 0, len(nums)-k+1)
	deque := make([]int, 0, k) // 存储下标

	for i := 0; i < len(nums); i++ {
		// 移除过期元素
		if len(deque) > 0 && deque[0] <= i-k {
			deque = deque[1:]
		}

		// 移除所有比当前元素小的元素
		for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[i] {
			deque = deque[:len(deque)-1]
		}

		deque = append(deque, i)

		// 当窗口形成后，记录最大值
		if i >= k-1 {
			result = append(result, nums[deque[0]])
		}
	}

	return result
}
