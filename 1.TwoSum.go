// Problem: 1. Two Sum
// Pattern: Map
// Link: https://leetcode.com/problems/two-sum/description/

package main

func twoSum(nums []int, target int) []int {

	numsFound := make(map[int]int, len(nums))

	for j, n := range nums {
		dv := target - n
		if i, exists := numsFound[dv]; exists {
			return []int{i, j}
		}
		numsFound[n] = j
	}

	return []int{}
}
