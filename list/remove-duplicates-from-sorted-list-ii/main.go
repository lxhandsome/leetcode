package main

// https://leetcode.cn/problems/remove-duplicates-from-sorted-list-ii/
type ListNode struct {
	Val  int
	Next *ListNode
}

// 1 2 2
func deleteDuplicates(head *ListNode) *ListNode {
	current := head
	h := &ListNode{}
	t := h
	for current != nil {
		flag := false
		for p := current.Next; p != nil && p.Val == current.Val; p = p.Next {
			flag = true
			current.Next = p.Next
		}
		if !flag {
			t.Next = current
			t = current
		}
		current = current.Next
		// 必须要将t指针Next置为nil，否则t.next可能携带最后不需要添加到h中的数字
		t.Next = nil
	}
	return h.Next
}

func testCases() *ListNode {
	return &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: nil}}}
}

func main() {
	deleteDuplicates(testCases())
}
