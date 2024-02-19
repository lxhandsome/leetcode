package main

// https://leetcode.cn/problems/partition-list/description/
type ListNode struct {
	Val  int
	Next *ListNode
}

/*
思路：
排序： 排序优先使用并归排序
寻找中间节点： 使用快慢指针，每次慢指针跳一步，快指针跳两步（快指针为空或者next为空停止），slow就是慢指针
递归结束状态： 当前节点为空或者当前的next为空
合并
*/
func sortList(head *ListNode) *ListNode {
	return mergeSort(head)
}

func findMiddle(h *ListNode) *ListNode {
	if h == nil {
		return h
	}
	slow := h
	fast := h.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func mergeSortedList(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			tail.Next = l1
			tail = l1
			l1 = l1.Next
		} else {
			tail.Next = l2
			tail = l2
			l2 = l2.Next
		}
	}
	if l1 != nil {
		tail.Next = l1
	}
	if l2 != nil {
		tail.Next = l2
	}
	return dummy.Next
}

func mergeSort(h *ListNode) *ListNode {
	if h == nil || h.Next == nil {
		return h
	}
	mid := findMiddle(h)
	// 获取第二条指针
	tail := mid.Next
	// 将第一条指针和第二条断开
	mid.Next = nil
	left := mergeSort(h)
	right := mergeSort(tail)
	return mergeSortedList(left, right)
}

func testCases() *ListNode {
	return &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 5, Next: &ListNode{Val: 2}}}}}}
}

func main() {
	sortList(testCases())
}
