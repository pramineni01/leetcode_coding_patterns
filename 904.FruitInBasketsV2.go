// Problem: 904. FruitInBaskets
// Pattern: Dynamic sliding window
// Link: https://leetcode.com/problems/fruit-into-baskets/description/

package main

import (
	"fmt"
	"math"
)

func main() {
	inp := []int{1, 2, 1}
	// inp := []int{0, 1, 2, 2}
	// inp := []int{1, 2, 3, 2, 2}

	out := findMaxFruitsPicked(inp)
	fmt.Println("Max fruits picked : ", out)
}

func findMaxFruitsPicked(inp []int) int {
	max := float64(0)

	type fruit_type int

	var start int
	var end int

	baskets := make(map[fruit_type]bool)

	for end = 0; end < len(inp); end++ {
		if (len(baskets) < 2) || (baskets[fruit_type(inp[end])]) {
			baskets[fruit_type(inp[end])] = true
		} else {
			max = math.Max(max, float64(end-start+1))
			delete(baskets, fruit_type(start))
			start++
		}
	}
	fmt.Printf("Start: %v, End: %v ", start, end)

	max = math.Max(max, float64(end-start+1))
	return int(max)
}
