// Problem: 1. Two Sum
// Pattern: Map
// Link: https://leetcode.com/problems/two-sum/description/

package main

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {

	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "1",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			name:   "2",
			nums:   []int{3, 2, 4},
			target: 6,
			want:   []int{1, 2},
		},
		{
			name:   "3",
			nums:   []int{3, 3},
			target: 6,
			want:   []int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.nums, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
