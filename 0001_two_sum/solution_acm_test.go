package main

import (
	"testing"
)

// 模拟ACM输入格式的测试
func TestTwoSumACM(t *testing.T) {
	// 这里可以添加针对ACM模式的特定测试
	// 由于main函数处理输入输出，我们可以测试twoSum函数
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{
			name:     "示例1",
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			name:     "示例2",
			nums:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		{
			name:     "示例3",
			nums:     []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := twoSum(tt.nums, tt.target)
			
			// 检查结果是否包含正确的下标（顺序可能不同）
			if len(result) != 2 {
				t.Errorf("Expected 2 elements, got %d", len(result))
				return
			}
			
			// 验证两数之和确实等于目标值
			if tt.nums[result[0]]+tt.nums[result[1]] != tt.target {
				t.Errorf("twoSum(%v, %d) = %v, but nums[%d] + nums[%d] = %d != %d", 
					tt.nums, tt.target, result, result[0], result[1], 
					tt.nums[result[0]]+tt.nums[result[1]], tt.target)
			}
			
			// 验证下标不相同
			if result[0] == result[1] {
				t.Errorf("Indices should be different, got [%d, %d]", result[0], result[1])
			}
		})
	}
}