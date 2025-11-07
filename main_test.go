package main

import (
	"reflect"
	"sort"
	"testing"
)

// Вспомогательная функция для сортировки результата (порядок троек не важен)
func sortResult(result [][]int) {
	for _, triple := range result {
		sort.Ints(triple)
	}
	sort.Slice(result, func(i, j int) bool {
		// Сортируем тройки для стабильного сравнения
		if result[i][0] != result[j][0] {
			return result[i][0] < result[j][0]
		}
		if result[i][1] != result[j][1] {
			return result[i][1] < result[j][1]
		}
		return result[i][2] < result[j][2]
	})
}

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		{
			name:     "example from leetcode",
			nums:     []int{-1, 0, 1, 2, -1, -4},
			expected: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name:     "empty array",
			nums:     []int{},
			expected: [][]int{},
		},
		{
			name:     "no triplets",
			nums:     []int{1, 2, 3, 4},
			expected: [][]int{},
		},
		{
			name:     "two sums on one nums[i]",
			nums:     []int{-100, 50, 50, 70, 30},
			expected: [][]int{{-100, 50, 50}, {-100, 30, 70}},
		},
		{
			name:     "repeating nums",
			nums:     []int{-100, 50, 50, 50, 50, 70, 30},
			expected: [][]int{{-100, 50, 50}, {-100, 30, 70}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := threeSum(tt.nums)

			sortResult(result)
			sortResult(tt.expected)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("threeSum(%v) = %v, expected %v", tt.nums, result, tt.expected)
			}
		})
	}
}
