package mergeintervals

import "sort"

// 题目：56. 合并区间
// 链接：https://leetcode.cn/problems/merge-intervals/
// 难度：Medium

// 思路：排序 + 贪心
// 1. 按区间起点排序
// 2. 遍历区间，如果当前区间与上一个区间重叠，则合并
// 3. 否则直接加入结果

// 时间复杂度：O(n log n)
// 空间复杂度：O(n)

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	// 按起点排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		last := result[len(result)-1]
		curr := intervals[i]

		if curr[0] <= last[1] {
			// 有重叠，合并
			if curr[1] > last[1] {
				last[1] = curr[1]
			}
		} else {
			// 无重叠，直接加入
			result = append(result, curr)
		}
	}

	return result
}
