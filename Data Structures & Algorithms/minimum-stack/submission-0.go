type MinStack struct {
	stack []int 
	min []int
}

func Constructor() MinStack {
	return MinStack{
		stack: []int{},
		min: []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)

	if len(this.min) == 0 {
		this.min = append(this.min, val)
	} else {
		this.min = append(this.min, minimum(val, this.min[len(this.min)-1]))
	}


}

func (this *MinStack) Pop() {
	this.stack = this.stack[0:len(this.stack)-1]
	this.min = this.min[0:len(this.min)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}

func minimum(a, b int) int {
	if a < b {
		return a
	}
	return b
}
