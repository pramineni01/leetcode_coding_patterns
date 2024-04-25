// Problem: 852. Peak Index in a Mountain Array
// Pattern: DFS + Two Pointers
// Link: https://leetcode.com/problems/peak-index-in-a-mountain-array/description/

package main

func peakIndexInMountainArray(arr []int) int {
	l, r := 0, len(arr)-1

	var mid int
	for l < r {
		mid = int((l + r) / 2)

		if arr[mid] < arr[mid+1] {
			l = mid + 1
		} else if arr[mid] > arr[mid+1] {
			r = mid
		}
	}
	return r
}
