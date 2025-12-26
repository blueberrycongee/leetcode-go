package main

import "testing"

// 模拟ACM输入格式的测试
func TestAddTwoNumbersACM(t *testing.T) {
	tests := []struct {
		name     string
		l1       []int
		l2       []int
		expected []int
	}{
		{
			name:     "示例1: [2,4,3] + [5,6,4] = [7,0,8]",
			l1:       []int{2, 4, 3},
			l2:       []int{5, 6, 4},
			expected: []int{7, 0, 8},
		},
		{
			name:     "示例2: [0] + [0] = [0]",
			l1:       []int{0},
			l2:       []int{0},
			expected: []int{0},
		},
		{
			name:     "示例3: [9,9,9,9,9,9,9] + [9,9,9,9] = [8,9,9,9,0,0,0,1]",
			l1:       []int{9, 9, 9, 9, 9, 9, 9},
			l2:       []int{9, 9, 9, 9},
			expected: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list1 := buildList(tt.l1)
			list2 := buildList(tt.l2)
			
			result := addTwoNumbers(list1, list2)
			
			// 将结果链表转换为数组以便比较
			resultArray := listToArray(result)
			
			if len(resultArray) != len(tt.expected) {
				t.Errorf("Length mismatch: got %v, want %v", resultArray, tt.expected)
				return
			}
			
			for i := 0; i < len(resultArray); i++ {
				if resultArray[i] != tt.expected[i] {
					t.Errorf("Element mismatch at index %d: got %v, want %v", i, resultArray, tt.expected)
					return
				}
			}
		})
	}
}

// 辅助函数：将链表转换为数组
func listToArray(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}