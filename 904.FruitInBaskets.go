// Problem: 904. FruitInBaskets
// Pattern: Dynamic sliding window
// Link: https://leetcode.com/problems/fruit-into-baskets/description/

package main

import (
	"math"
)

// func main() {
// 	// inp := []int{1, 2, 1}
// 	// inp := []int{0, 1, 2, 2}
// 	inp := []int{1, 2, 3, 2, 2}

// 	out := findMaxFruitsPicked(inp)
// 	fmt.Println("Max fruits picked : ", out)
// }

func findMaxFruitsPicked(inp []int) int {
	max := 0

	windowPickedCount := 0
	start := 0

	baskets := make(map[int]int)
	for end := 0; end < len(inp); end++ {

		if len(baskets) < 2 {
			baskets[inp[end]] += 1
			windowPickedCount++
			continue
		}

		if _, exists := baskets[inp[end]]; exists {
			baskets[inp[end]]++
			windowPickedCount++
		} else {
			max = int(math.Max(float64(max), float64(windowPickedCount)))
			windowPickedCount -= baskets[inp[start]]
			delete(baskets, inp[start])
			start++

			baskets[inp[end]] += 1
			windowPickedCount++
		}
	}

	max = int(math.Max(float64(max), float64(windowPickedCount)))
	return max
}
