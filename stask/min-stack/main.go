package main

/*
思路：
使用连个栈保存数据，
最小栈保存当前的最小值
普通栈保存数据
*/

// https://leetcode.cn/problems/min-stack/description/

type MinStack struct {
	min   []int
	stack []int
}

func Constructor() MinStack {
	return MinStack{
		min:   make([]int, 0),
		stack: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	min := this.GetMin()
	if val < min {
		this.min = append(this.min, val)
	} else {
		this.min = append(this.min, min)
	}
	this.stack = append(this.stack, val)
}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	this.min = this.min[:len(this.min)-1]
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return 0
	}
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.min) == 0 {
		return 1 << 31
	}
	min := this.min[len(this.min)-1]
	return min
}

// type MinStack struct {
//     min []int
//     stack []int
// }

// /** initialize your data structure here. */
// func Constructor() MinStack {
//     return MinStack{
//         min: make([]int, 0),
//         stack: make([]int, 0),
//     }
// }

// func (this *MinStack) Push(x int)  {
//     min := this.GetMin()
//     if x < min {
//         this.min = append(this.min, x)
//     } else {
//         this.min = append(this.min, min)
//     }
//     this.stack = append(this.stack, x)
// }

// func (this *MinStack) Pop()  {
//     if len(this.stack) == 0 {
//         return
//     }
//     this.stack = this.stack[:len(this.stack)-1]
//     this.min = this.min[:len(this.min)-1]
// }

// func (this *MinStack) Top() int {
//     if len(this.stack) == 0 {
//         return 0
//     }
//     return this.stack[len(this.stack)-1]
// }

// func (this *MinStack) GetMin() int {
//     if len(this.min) == 0 {
//         return 1 << 31
//     }
//     min := this.min[len(this.min)-1]
//     return min
// }
