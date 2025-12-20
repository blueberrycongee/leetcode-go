package movezeroes

// 题目：283. 移动零
// 链接：https://leetcode.cn/problems/move-zeroes/
// 难度：Easy
// 描述：给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
// 请注意 ，必须在不复制数组的情况下原地对数组进行操作。

func moveZeroes(nums []int) {
	// 双指针法
	// slow: 指向当前已处理好的非零序列的尾部（下一个非零元素该放的位置）
	// fast: 用于遍历数组查找非零元素
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		// 如果发现非零元素
		if nums[fast] != 0 {
			// 将它交换到 slow 的位置
			nums[slow], nums[fast] = nums[fast], nums[slow]
			// slow 指针前进一步，准备接下一个非零元素
			slow++
		}
	}
}
