// Problem: 11. Container With Most Water
// Pattern: Two Pointer
// Link: https://leetcode.com/problems/container-with-most-water/description/

package main

import "math"

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	maxVol := float64(0)

	for l < r {
		vol := math.Min(float64(height[l]), float64(height[r])) * float64(r-l)
		maxVol = math.Max(maxVol, vol)
	}

	return int(maxVol)
}
