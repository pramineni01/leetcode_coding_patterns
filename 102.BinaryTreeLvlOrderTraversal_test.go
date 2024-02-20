// Problem: 102. Binary Tree Level Order Traversal
// Pattern: Breadth First Search (BFS)
// Link: https://leetcode.com/problems/binary-tree-level-order-traversal/description/

package main

import (
	"reflect"
	"testing"
)

func Test_levelOrder(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1",
			root: makeTree([]int{3,9,20,-1, -1,15,7})
			want: [][]int{[]int{3},[]int{9,20},[]int{15,7}},
		},
		{
			name: "1",
			root: makeTree([]int{3,9,20,-1, -1,15,7})
			want: [][]int{[]int{3},[]int{9,20},[]int{15,7}},
		}
		[3,9,20,null,null,15,7]

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
