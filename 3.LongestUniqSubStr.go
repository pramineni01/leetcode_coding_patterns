// Problem: 3. Longest Substring Without Repeating Characters
// Pattern: Dynamic sliding window
// Link: https://leetcode.com/problems/longest-substring-without-repeating-characters/description/

package main

import (
	"math"
)

// func main() {
// 	// inp := "abcabcbb" //Output: 3
// 	// inp := "bbbbb" // Output: 1
// 	// inp := "pwwkew" //Output: 3
// 	inp := "abba" //Output: 2
// 	// inp := " " //Output: 1

// 	out := findLongestSubStringWithUniqChars(inp)
// 	fmt.Println("Longest substring with unique characters: ", out)
// }

func findLongestSubStringWithUniqChars(s string) int {

	var length float64
	start := float64(-1)

	charSet := make(map[rune]int)
	for end := 0; end < len(s); end++ {
		if idx, exists := charSet[rune(s[end])]; !exists {
			charSet[rune(s[end])] = end
		} else {
			start = math.Max(start, float64(idx))
			charSet[rune(s[end])] = end
		}
		length = math.Max(length, float64(end)-start)
	}

	return int(length)
}
