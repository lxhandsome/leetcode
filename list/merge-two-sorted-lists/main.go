package main

// https://leetcode.cn/problems/merge-two-sorted-lists/
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	h := &ListNode{}
	tail := h
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			tmp := list1.Next
			list1.Next = nil
			tail.Next = list1
			tail = list1
			list1 = tmp
		} else {
			tmp := list2.Next
			list2.Next = nil
			tail.Next = list2
			tail = list2
			list2 = tmp
		}
	}
	if list1 != nil {
		tail.Next = list1
	}
	if list2 != nil {
		tail.Next = list2
	}
	return h.Next
}

func testCases() (*ListNode, *ListNode) {
	return &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
		&ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}
}

func main() {
	mergeTwoLists(testCases())
}
