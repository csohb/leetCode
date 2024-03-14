type MinStack struct {
    stack []int
    minStack []int
}


func Constructor() MinStack {
    return MinStack{
        stack : make([]int, 0, 30000),
        minStack : make([]int, 0, 30000),
    }
}


func (this *MinStack) Push(val int)  {
    min := val
    if len(this.minStack) > 0 && this.minStack[len(this.minStack) - 1] < min {
        min = this.minStack[len(this.minStack) -1 ]
    }
    this.stack = append(this.stack, val)
    this.minStack = append(this.minStack, min)
}


func (this *MinStack) Pop()  {
    this.stack = this.stack[:len(this.stack)-1]
}


func (this *MinStack) Top() int {
    return this.stack[len(this.stack)-1]
}


func (this *MinStack) GetMin() int {
    return this.minStack[len(this.minStack)-1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */