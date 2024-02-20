// Problem: 637. Average of Levels in Binary Tree
// Pattern: BFS
// Link: https://leetcode.com/problems/average-of-levels-in-binary-tree/description/

package main

// * Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {

	result := []float64{}
	if root == nil {
		return result
	}

	q := []*TreeNode{root}

	for len(q) > 0 {
		currLvlLen := len(q)

		currLvlSum := float64(0)
		for i := 0; i < currLvlLen; i++ {
			currLvlSum += float64(q[i].Val)

			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}

		}

		result = append(result, float64(currLvlSum)/float64(currLvlLen))

		// deque current level
		q = q[currLvlLen:]

	}
	return result
}
