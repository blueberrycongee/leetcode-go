package twosum

// 题目：1. 两数之和
// 链接：https://leetcode.cn/problems/two-sum/
// 难度：Easy

// 思路：哈希表
// 1. 遍历数组，对于每个数num，查找 target-num 是否在map中
// 2. 如果在，返回两个下标；如果不在，把num存入map

// 时间复杂度：O(n)
// 空间复杂度：O(n)

func twoSum(nums []int, target int) []int {
	// m: 数值 -> 下标
	m := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if j, ok := m[complement]; ok {
			return []int{j, i}
		}
		m[num] = i
	}

	return nil
}
