package main

// https://leetcode.cn/problems/partition-list/description/
type ListNode struct {
	Val  int
	Next *ListNode
}

// 思路: 新建两个头结点,一个保存所有小于x的节点,一个保存所有大于等于x的节点
// 对链表进行一次遍历,根据节点的大小使用尾插法插入不同的链表中,这样能够保存链表的顺序
// 最后对链表进行拼接
func partition(head *ListNode, x int) *ListNode {
	less := &ListNode{}
	greater := &ListNode{}
	lt := less
	gt := greater
	for head != nil {
		tmp := head.Next
		if head.Val < x {
			lt.Next = head
			lt = head
			lt.Next = nil
		} else {
			gt.Next = head
			gt = head
			gt.Next = nil
		}
		head = tmp
	}
	if less.Next == nil {
		return greater.Next
	}
	lt.Next = greater.Next
	return less.Next
}

func testCases() *ListNode {
	return &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 5, Next: &ListNode{Val: 2}}}}}}
}

func main() {
	partition(testCases(), 3)
}
