// Problem: 704. Binary Search
// Pattern: DFS
// Link: https://leetcode.com/problems/binary-search/description/

package main

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := (l + r) / 2

		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}
