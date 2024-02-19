package main

// https://leetcode.cn/problems/remove-duplicates-from-sorted-list/
type ListNode struct {
	Val  int
	Next *ListNode
}

// 1 1 1 2 3 5
// 如果链表无需的话需要 每个元素扫描完剩余的其他元素 n方的复杂度
// 现在链表有序，只需要遍历一遍，以第一个数为基准，遍历如果后一个数与当前相同的话，将直接把后一个数丢弃， 如果不相同的话更新current为current.next
func deleteDuplicates(head *ListNode) *ListNode {
	current := head
	for current != nil {
		for p := current.Next; p != nil && p.Val == current.Val; p = p.Next {
			current.Next = p.Next
		}
		current = current.Next
	}
	return head
}
