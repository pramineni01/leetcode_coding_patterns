// Problem: 852. Peak Index in a Mountain Array
// Pattern: DFS + Two Pointers
// Link: https://leetcode.com/problems/peak-index-in-a-mountain-array/description/

package main

import "testing"

func Test_peakIndexInMountainArray(t *testing.T) {

	tests := []struct {
		name string
		nums []int
		want int
	}{
		// {
		// 	name: "1",
		// 	nums: []int{0, 1, 0},
		// 	want: 1,
		// },
		{
			name: "2",
			nums: []int{0, 2, 1, 0},
			want: 1,
		},
		{
			name: "3",
			nums: []int{0, 10, 5, 2},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := peakIndexInMountainArray(tt.nums); got != tt.want {
				t.Errorf("peakIndexInMountainArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
