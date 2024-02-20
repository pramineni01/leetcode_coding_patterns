// Problem: 21. Merge Two Sorted Lists
// Pattern: Two Pointer
// Link: https://leetcode.com/problems/merge-two-sorted-lists/description/

package main

// * Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	list3 := &ListNode{Val: -1}
	head := list3

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			list3.Next = &ListNode{Val: list1.Val}
			list3, list1 = list3.Next, list1.Next
		} else if list1.Val > list2.Val {
			list3.Next = &ListNode{Val: list2.Val}
			list3, list2 = list3.Next, list2.Next
		} else { // both equal
			list3.Next = &ListNode{Val: list1.Val}
			list3, list1 = list3.Next, list1.Next

			list3.Next = &ListNode{Val: list2.Val}
			list3, list2 = list3.Next, list2.Next
		}
	}

	for list1 != nil {
		list3.Next = &ListNode{Val: list1.Val}
		list3, list1 = list3.Next, list1.Next
	}

	for list2 != nil {
		list3.Next = &ListNode{Val: list2.Val}
		list3, list2 = list3.Next, list2.Next
	}

	return head.Next
}
