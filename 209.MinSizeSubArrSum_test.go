// Problem: 209. Minimum Size Subarray Sum
// Problem Pattern: Dynamic Sliding Window Pattern
// Link: https://leetcode.com/problems/minimum-size-subarray-sum/description/

package main

import "testing"

func Test_minSubArrayLen(t *testing.T) {
	tests := []struct {
		name   string
		target int
		nums   []int
		want   int
	}{
		{
			name:   "1",
			target: 7,
			nums:   []int{2, 3, 1, 2, 4, 3},
			want:   2,
		},
		{
			name:   "2",
			target: 4,
			nums:   []int{1, 4, 4},
			want:   1,
		},
		{
			name:   "3",
			target: 11,
			nums:   []int{1, 1, 1, 1, 1, 1, 1, 1},
			want:   0,
		},
		{
			name:   "4",
			target: 11,
			nums:   []int{1, 2, 3, 4, 5},
			want:   3,
		},
		{
			name:   "15",
			target: 15,
			nums:   []int{5, 1, 3, 5, 10, 7, 4, 9, 2, 8},
			want:   2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubArrayLen(tt.target, tt.nums); got != tt.want {
				t.Errorf("minSubArrayLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
