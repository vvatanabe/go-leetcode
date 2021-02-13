package main

type MyStack struct {
	q1 []int
	q2 []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		q1: []int{},
		q2: []int{},
	}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.q1 = append(this.q1, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	size := len(this.q1)
	for i := 0; i < size-1; i++ {
		this.q2 = append(this.q2, this.q1[0])
		this.q1 = this.q1[1:]
	}
	top := this.q1[0]
	this.q1 = this.q2
	this.q2 = nil
	return top
}

/** Get the top element. */
func (this *MyStack) Top() int {
	top := this.Pop()
	this.q1 = append(this.q1, top)
	return top
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.q1) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
