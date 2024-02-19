// Problem: 209. Minimum Size Subarray Sum
// Problem Pattern: Dynamic Sliding Window Pattern
// Link: https://leetcode.com/problems/minimum-size-subarray-sum/description/

package main

import "math"

func minSubArrayLenV2(target int, nums []int) int {
	l, r := 0, 0
	minLen := math.MaxInt32

	sum := 0
	for r < len(nums) {
		sum += nums[r]

		for (sum >= target) && (l < r+1) {
			if sum >= target {
				minLen = min(minLen, r-l+1)
			}
			sum -= nums[l]
			l++
		}

		r++
	}

	if minLen == math.MaxInt32 {
		return 0
	}
	return minLen
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
