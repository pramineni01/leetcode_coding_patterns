// Problem: 21. Merge Two Sorted Lists
// Pattern: Two Pointer
// Link: https://leetcode.com/problems/merge-two-sorted-lists/description/

package main

import (
	"reflect"
	"testing"
)

func makeList(inp []int) *ListNode {
	if len(inp) == 0 {
		return nil
	}

	list := &ListNode{Val: inp[0]}
	head := list
	for _, v := range inp[1:] {
		list.Next = &ListNode{Val: v}
		list = list.Next
	}
	return head
}

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		list1 *ListNode
		list2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "1",
			args: args{
				list1: makeList([]int{1, 2, 4}),
				list2: makeList([]int{1, 3, 4}),
			},
			want: makeList([]int{1, 1, 2, 3, 4, 4}),
		},
		{
			name: "2",
			args: args{
				list1: makeList([]int{}),
				list2: makeList([]int{}),
			},
			want: makeList([]int{}),
		},
		{
			name: "3",
			args: args{
				list1: makeList([]int{}),
				list2: makeList([]int{0}),
			},
			want: makeList([]int{0}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.list1, tt.args.list2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
