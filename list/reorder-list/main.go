package main

// https://leetcode.cn/problems/reorder-list/description/
type ListNode struct {
	Val  int
	Next *ListNode
}

/*
思路：
1、先找到链表的中间节点，将链表分成两节
2、将后面一节进行翻转
3、将第一节链表与翻转后的第二节链表进行合并，从第一节开始
*/
func reorderList(head *ListNode) {
	mid := findMiddle(head)
	next := mid.Next
	mid.Next = nil
	next = reverseList(next)
	head = mergeTwoList(head, next)
}

func findMiddle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func mergeTwoList(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	flag := true
	for l1 != nil && l2 != nil {
		if flag {
			tail.Next = l1
			tail = tail.Next
			l1 = l1.Next
			flag = !flag
			continue
		}
		tail.Next = l2
		tail = tail.Next
		l2 = l2.Next
		flag = !flag
	}
	if l1 != nil {
		tail.Next = l1
	}
	if l2 != nil {
		tail.Next = l2
	}
	return dummy.Next
}

func reverseList(head *ListNode) *ListNode {
	var p *ListNode
	for head != nil {
		tmp := head.Next
		head.Next = p
		p = head
		head = tmp
	}
	return p
}

func testCases() *ListNode {
	return &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
}

func testCases1() *ListNode {
	return &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}
}

func main() {
	reorderList(testCases())
}
