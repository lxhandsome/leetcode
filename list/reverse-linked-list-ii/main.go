package main

// https://leetcode.cn/problems/reverse-linked-list-ii/description/
type ListNode struct {
	Val  int
	Next *ListNode
}

// pre : 指向翻转范围的前一个节点
// head: 正向遍历链表,当前节点
// next: l - r翻转之后的链表
// result:  pre -> next -> head
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return nil
	}
	// 因为有可能从头开发翻转，导致返回
	h := &ListNode{}
	h.Next = head
	head = h
	var pre *ListNode
	i := 0
	for i < left {
		pre = head
		head = head.Next
		i++
	}
	// 还需要记录 next tail节点
	var next *ListNode
	// 翻转选中的部分
	// 第一个节点就是翻转后的最后一个节点
	nextTail := head
	for head != nil && i <= right {
		tmp := head.Next
		head.Next = next
		next = head
		head = tmp
		i++
	}
	pre.Next = next
	nextTail.Next = head
	return h.Next
}

func testCases() *ListNode {
	return &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
}

func main() {
	reverseBetween(testCases(), 2, 4)
}
