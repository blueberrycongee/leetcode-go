package longestconsecutive

// 题目：128. 最长连续序列
// 链接：https://leetcode.cn/problems/longest-consecutive-sequence/
// 难度：Medium

// 思路：哈希表
// 1. 把所有数存入 set
// 2. 遍历每个数，如果 num-1 不存在，说明 num 是序列起点
// 3. 从起点往后数，统计长度

func longestConsecutive(nums []int) int {
	// 1. 把所有数存入 set（Go用map代替set）
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	maxLen := 0

	// 2. 遍历每个数
	for num := range numSet {
		// 关键：如果 num-1 存在，说明 num 不是起点，跳过
		// 这样保证每个序列只从起点数一次，O(n)
		if numSet[num-1] {
			continue
		}

		// num 是起点，往后数
		currentNum := num
		currentLen := 1

		// 一直数到断开为止
		for numSet[currentNum+1] {
			currentNum++
			currentLen++
		}

		// 更新最大长度
		if currentLen > maxLen {
			maxLen = currentLen
		}
	}

	return maxLen
}
