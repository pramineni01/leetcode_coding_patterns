// Problem: 167. Two Sum II - Input Array Is Sorted
// Pattern: Two Pointers
// Link: https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/

package main

import (
	"reflect"
	"testing"
)

func Test_twoSumSorted(t *testing.T) {

	tests := []struct {
		name    string
		numbers []int
		target  int
		want    []int
	}{
		{
			name:    "1",
			numbers: []int{2, 7, 11, 15},
			target:  9,
			want:    []int{1, 2},
		},
		{
			name:    "2",
			numbers: []int{2, 3, 4},
			target:  6,
			want:    []int{1, 3},
		},
		{
			name:    "3",
			numbers: []int{-1, 0},
			target:  -1,
			want:    []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.numbers, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
