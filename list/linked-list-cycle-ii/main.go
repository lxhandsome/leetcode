package main

// https://leetcode.cn/problems/linked-list-cycle-ii/description/
type ListNode struct {
	Val  int
	Next *ListNode
}

/*
思路: 假设从头结点到环入口距离为a，环长度为b
使用快慢指针，fast快指针每次奏两步，slow慢指针每次走一步，当快慢指针第一次相遇的时候
f  = a + k* b
s = a + j * b
推理可得: f = n * b

fast每次走的步数是slow的二倍当相遇后：f = 2 *s

综上： f = 2nb ，s = nb

如果从头结点到环入口处走过的距离 d = a + nb

现在要求环入口节点相当于从头结点 出发再走d部，现在s已经走了nb部只需要再走a步即可到达入口点，现在a不知道是多少，但是从头结点到环入口的距离就是a
相当于s先行nb后，此时s f相遇，再次相遇后将f指向头结点，再次行走a步，f到达入环点，同时s也行走a步也到达入环点此时f s再次相遇相遇点即为入环点


*/

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow != fast {
			continue
		}
		fast = head
		for slow != fast {
			slow = slow.Next
			fast = fast.Next
		}
		return slow
	}
	return nil
}
