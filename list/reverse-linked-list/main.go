package main

// https://leetcode.cn/problems/reverse-linked-list/
type ListNode struct {
	Val  int
	Next *ListNode
}

func testCases() *ListNode {
	return &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}}
}

// 头插法
// 在进行头插的过程中会有head.next = tmp 操作会使得原先的链表断开，所以需要使用next保存当前节点的next，最后在赋值回来，保证能够遍历到下一个节点
func reverseList1(head *ListNode) *ListNode {
	h := &ListNode{}
	for head != nil {
		next := head.Next
		tmp := h.Next
		h.Next = head
		head.Next = tmp
		head = next
	}
	return h.Next
}

// 遍历链表的时候记录前一个节点, 不需要new 一个head节点了
// 遍历正序链表,将逆序
func reverseList(head *ListNode) *ListNode {
	var p *ListNode
	for head != nil {
		// 记录下一个节点
		next := head.Next
		head.Next = p
		// 将逆序之后的链表赋值给p
		p = head
		// 接着遍历原先剩余的节点
		head = next
	}
	return head
}

func main() {
	reverseList(testCases())
}
