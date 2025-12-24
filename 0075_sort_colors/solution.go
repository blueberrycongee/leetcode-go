package sortcolors

// 题目：75. 颜色分类
// 链接：https://leetcode.cn/problems/sort-colors/
// 难度：Medium

// 思路：三指针（荷兰国旗问题）
// 1. p0 指向下一个放 0 的位置，p2 指向下一个放 2 的位置
// 2. curr 遍历数组：
//    - 遇到 0：与 p0 交换，p0++, curr++
//    - 遇到 2：与 p2 交换，p2--（curr不动，因为换来的数还需检查）
//    - 遇到 1：curr++

// 时间复杂度：O(n)
// 空间复杂度：O(1)

func sortColors(nums []int) {
	p0, curr, p2 := 0, 0, len(nums)-1

	for curr <= p2 {
		if nums[curr] == 0 {
			nums[p0], nums[curr] = nums[curr], nums[p0]
			p0++
			curr++
		} else if nums[curr] == 2 {
			nums[p2], nums[curr] = nums[curr], nums[p2]
			p2--
		} else {
			curr++
		}
	}
}
