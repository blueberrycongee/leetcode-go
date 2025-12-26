package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ==================== 题目信息 ====================
// 题目：1. 两数之和 (ACM模式)
// 难度：Easy
// 链接：https://leetcode.cn/problems/two-sum/
//
// ==================== 输入格式 ====================
// 第一行：两个整数 n 和 target，用空格分隔
//   - n: 数组长度 (2 <= n <= 10^4)
//   - target: 目标值
// 第二行：n 个整数，用空格分隔，表示数组元素
//
// ==================== 输出格式 ====================
// 一行，包含两个整数，用空格分隔
// 表示两个数的下标（从0开始），使得它们的和等于target
//
// ==================== 输入输出示例 ====================
// 【示例1】
// 输入：
// 4 9
// 2 7 11 15
// 输出：
// 0 1
// 解释：nums[0] + nums[1] = 2 + 7 = 9
//
// 【示例2】
// 输入：
// 3 6
// 3 2 4
// 输出：
// 1 2
// 解释：nums[1] + nums[2] = 2 + 4 = 6
//
// 【示例3】
// 输入：
// 2 6
// 3 3
// 输出：
// 0 1
// 解释：nums[0] + nums[1] = 3 + 3 = 6
//
// ==================== 解题思路 ====================
// 思路：哈希表
// 1. 遍历数组，对于每个数num，查找 target-num 是否在map中
// 2. 如果在，输出两个下标；如果不在，把num存入map
//
// 时间复杂度：O(n)
// 空间复杂度：O(n)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	
	// 读取第一行：n 和 target
	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	n, _ := strconv.Atoi(parts[0])
	target, _ := strconv.Atoi(parts[1])
	
	// 读取第二行：数组元素
	scanner.Scan()
	numsStr := strings.Fields(scanner.Text())
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i], _ = strconv.Atoi(numsStr[i])
	}
	
	// 使用哈希表求解
	result := twoSum(nums, target)
	
	// 输出结果
	fmt.Printf("%d %d\n", result[0], result[1])
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int) // 数值 -> 下标

	for i, num := range nums {
		complement := target - num

		if j, ok := m[complement]; ok {
			return []int{j, i}
		}

		m[num] = i
	}

	return nil
}