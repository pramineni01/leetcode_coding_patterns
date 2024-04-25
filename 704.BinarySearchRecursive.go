// Problem: 704. Binary Search - Recursive variant
// Pattern: DFS
// Link: https://leetcode.com/problems/binary-search/description/

package main

func binarySearch(nums []int, l, r, target int) int {
	if l > r {
		return -1
	}

	mid := int((l + r) / 2)

	if nums[mid] == target {
		return mid
	}
	if nums[mid] < target {
		return binarySearch(nums, mid+1, r, target)
	} else {
		return binarySearch(nums, l, mid-1, target)
	}
}

func recursiveSearch(nums []int, target int) int {
	return binarySearch(nums, 0, len(nums)-1, target)
}
