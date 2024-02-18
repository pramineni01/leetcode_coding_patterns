// Problem: 643. Maximum Average Subarray I
// Problem Pattern: Sliding Window Pattern (following code using naive approach)
// Link: https://leetcode.com/problems/maximum-average-subarray-i/description/

package main

import (
	"fmt"
	"math"
)

func main() {

	inp := []int{1, 12, -5, -6, 50, 3}
	k := 4

	out := findMaxAvgSubArray(inp, k)
	fmt.Println("Output: %v", out)
}

func findMaxAvgSubArray(inp []int, k int) float64 {
	max := float64(0)

	for i := 0; i < len(inp); i++ {
		sum := 0
		for j := 0; j < k; j++ {
			if !(i+(k-1) > len(inp)-1) {
				sum += inp[i+j]
			} else {
				i = len(inp)
				j = k
			}
		}
		if avg := float64(sum) / float64(k); avg > 0 {
			max = math.Max(max, avg)
		}
	}

	return max
}
