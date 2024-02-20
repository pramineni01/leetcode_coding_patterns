// // Problem: 643. Maximum Average Subarray I
// // Problem Pattern: Sliding Window Pattern (following code using naive approach)
// // Link: https://leetcode.com/problems/maximum-average-subarray-i/description/

package main

import (
	"math"
)

// func main() {
// 	inp := []int{1, 12, -5, -6, 50, 3}
// 	k := 4

// 	out := findMaxAvgSubArraySlidingWindow(inp, k)
// 	fmt.Println("Output: ", out)
// }

func findMaxAvgSubArraySlidingWindow(inp []int, k int) float64 {
	max := float64(0)

	windowSum := 0
	start := 0

	for end := 0; end < len(inp); end++ {
		windowSum += end

		if (end - start + 1) == k {
			max = math.Max(max, float64(windowSum)/float64(k))

			windowSum -= inp[start]
			start++
		}
	}

	return max
}
