// Problem: 704. Binary Search - Recursive variant
// Pattern: DFS
// Link: https://leetcode.com/problems/binary-search/description/

package main

import "testing"

func Test_recursiveSearch(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				nums:   []int{-1, 0, 3, 5, 9, 12},
				target: 9,
			},
			want: 4,
		},
		{
			name: "2",
			args: args{
				nums:   []int{-1, 0, 3, 5, 9, 12},
				target: 2,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := recursiveSearch(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
