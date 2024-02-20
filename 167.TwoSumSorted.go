// Problem: 167. Two Sum II - Input Array Is Sorted
// Pattern: Two Pointers
// Link: https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/

package main

func twoSumSorted(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for {
		sum := numbers[l] + numbers[r]
		if sum < target {
			l++
		} else if sum > target {
			r--
		} else if sum == target {
			return []int{l + 1, r + 1}
		}
	}
	return []int{}
}
