package twosum

// 题目：1. 两数之和
// 链接：https://leetcode.cn/problems/two-sum/
// 难度：Easy

// 思路：哈希表
// 1. 遍历数组，对于每个数num，查找 target-num 是否在map中
// 2. 如果在，返回两个下标；如果不在，把num存入map

// 时间复杂度：O(n)
// 空间复杂度：O(n)

// func 函数名(参数名 参数类型) 返回值类型
// Go特点：类型写在变量后面，返回值写在最后
func twoSum(nums []int, target int) []int {

	// make() = 创建并初始化（用于slice/map/channel）
	// map[int]int = key是int, value是int
	// := 是声明+赋值，自动推断类型（等价于 var m map[int]int = make(...)）
	//
	// 为什么用map？O(1)时间查找，比遍历O(n)快
	m := make(map[int]int) // m: 数值 -> 下标

	// for i, num := range nums = 遍历数组
	//   - i = 下标 (0, 1, 2, ...)
	//   - num = 当前元素的值
	//   - range = 自动从头遍历到尾，不需要写循环条件
	//
	// 如果只要值不要下标：for _, num := range nums
	// 如果只要下标不要值：for i := range nums
	for i, num := range nums {
		complement := target - num // 需要找的另一个数

		// if 语句; 条件 { }  ← Go特殊语法，分号前先执行，分号后判断
		//
		// j, ok := m[complement]
		//   - Go的map取值返回两个东西！（一个操作返回多个值是Go特性）
		//   - j = value（complement对应的下标）
		//   - ok = true/false（key是否存在）
		//
		// 为什么要ok？因为m[不存在的key]不报错，返回0，无法区分"值是0"还是"不存在"
		if j, ok := m[complement]; ok {
			// []int{j, i} = 创建int切片，包含j和i
			// 注意是大括号{}，不是小括号()
			return []int{j, i}
		}

		// m[num] = i  ← 赋值用 =（不是:=，因为m已存在）
		// 把当前数存入map："数值num在下标i"
		// 存起来给后面的数找
		m[num] = i
	}

	// nil = 空/不存在，类似C的NULL、Python的None
	return nil
}
