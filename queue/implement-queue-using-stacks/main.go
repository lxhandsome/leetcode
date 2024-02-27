package main

type MyQueue struct {
	stack1 []int
	stack2 []int
}

func Constructor() MyQueue {
	return MyQueue{
		stack1: make([]int, 0),
		stack2: make([]int, 0),
	}

}

func (this *MyQueue) Push(x int) {
	this.stack1 = append(this.stack1, x)
}

func (this *MyQueue) Pop() int {
	res := this.peek()
	this.stack2 = this.stack2[:len(this.stack2)-1]
	return res
}

func (this *MyQueue) peek() int {
	if this.Empty() {
		return 0
	}
	if this.stack2IsEmpty() {
		this.mv2Stack2()
	}
	res := this.stack2[len(this.stack2)-1]
	return res
}

func (this *MyQueue) Peek() int {
	return this.peek()
}

func (this *MyQueue) stack2IsEmpty() bool {
	return len(this.stack2) == 0
}

func (this *MyQueue) stack1IsEmpty() bool {
	return len(this.stack1) == 0
}

func (this *MyQueue) mv2Stack2() {
	for i := len(this.stack1) - 1; i >= 0; i-- {
		this.stack2 = append(this.stack2, this.stack1[i])
	}
	this.stack1 = make([]int, 0)
}

func (this *MyQueue) Empty() bool {
	return this.stack2IsEmpty() && this.stack1IsEmpty()
}
