type MinStack struct {
	stack  []int
	mStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:  []int{},
		mStack: []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)

	if len(this.mStack) == 0 {
		this.mStack = append(this.mStack, val)
	} else {
		curmin := this.mStack[len(this.mStack)-1]
		if val < curmin {
			this.mStack = append(this.mStack, val)
		} else {
			this.mStack = append(this.mStack, curmin)
		}
	}
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.mStack = this.mStack[:len(this.mStack)-1]

}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.mStack[len(this.mStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */