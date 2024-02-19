// Problem: 209. Minimum Size Subarray Sum
// Problem Pattern: Dynamic Sliding Window Pattern
// Link: https://leetcode.com/problems/minimum-size-subarray-sum/description/

package main

import "fmt"

// func main() {
// 	inp := []int{2, 3, 1, 2, 4, 3}
// 	target := 10
// 	out := minSubArrayLen(target, inp)
// 	fmt.Println("Minimum size subarray sum: ", out)
// }

func addAll(nums []int) int {
	fmt.Printf("AddAll: %v\n", nums)
	var sum int
	for _, v := range nums {
		sum += v
	}
	return sum
}

func minSubArrayLen(target int, nums []int) int {

	min_start, min_end := -1, -1

	sum := 0
	for start, end := 0, 0; end < len(nums); end++ {
		sum += nums[end]
		if sum < target {
			continue
		} else {
			fmt.Printf("Start: %v, End: %v\n", start, end)

			if sum > target {
				// reduce window size at "start"
				i := start
				for ; (i <= end) && (sum-nums[i] >= target); i++ {
					sum -= nums[i]
					continue
				}

				if i > start {
					start = i
				}
			}

			if (min_end <= -1) || ((min_end - min_start) > (end - start)) {
				min_start, min_end = start, end
			}
		}
	}

	if min_start == -1 && min_end == -1 {
		return 0
	} else if min_start == -1 {
		return min_end + 1
	} else {
		fmt.Printf("Min_start: %v, Min_end: %v\n", min_start, min_end)
		return min_end - min_start + 1
	}
}
