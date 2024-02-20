// Problem: 102. Binary Tree Level Order Traversal
// Pattern: Breadth First Search (BFS)
// Link: https://leetcode.com/problems/binary-tree-level-order-traversal/description/

package main

// * Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Traverse tree and output list of lists representing each level
func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}

	if root == nil {
		return result
	}

	// enque root node
	q := []*TreeNode{root}

	for len(q) > 0 {
		length := len(q)
		levelNodes := []int{}

		for i := 0; i < length; i++ {
			levelNodes = append(levelNodes, q[i].Val)

			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}

			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
		}

		// deque node
		q = q[length:]
		result = append([][]int{levelNodes}, result...)
	}

	return result
}
