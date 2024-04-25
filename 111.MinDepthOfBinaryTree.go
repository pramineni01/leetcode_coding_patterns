// Problem: 111. Minimum Depth of Binary Tree
// Pattern: Breadth First Search (BFS)
// Link: https://leetcode.com/problems/minimum-depth-of-binary-tree/description/

package main

// * Definition for a binary tree node.
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func minDepth(root *TreeNode) int {
	minDepth := 0

	if root == nil {
		return minDepth
	}

	q := []*TreeNode{root}

	// currentDepth := 0
	for len(q) > 0 {
		minDepth++
		currLvlLen := len(q)
		for i := 0; i < currLvlLen; i++ {
			if q[i].Left == nil && q[i].Right == nil {
				return minDepth
			}
			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
		}
	}
	return minDepth
}
