package groupanagrams

import "slices"

// 题目：49. 字母异位词分组
// 链接：https://leetcode.cn/problems/group-anagrams/
// 难度：Medium

// 思路：哈希表
// 1. 把每个字符串排序，作为map的key
// 2. 相同key的字符串放到同一个数组里

func groupAnagrams(strs []string) [][]string {
	// 你来写！
	m := make(map[string][]string)
	for _, str := range strs {
		bytes := []byte(str)

		slices.Sort(bytes)

		key := string(bytes)
		m[key] = append(m[key], str)
	}

	res := make([][]string, 0)

	for _, v := range m {
		res = append(res, v)
	}
	return res
}
