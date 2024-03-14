type MinStack struct {
    stack []int
}


func Constructor() MinStack {
    return MinStack{}
}


func (this *MinStack) Push(val int)  {
    this.stack = append(this.stack, val)
}


func (this *MinStack) Pop()  {
    this.stack = this.stack[:len(this.stack)-1]
}


func (this *MinStack) Top() int {
    return this.stack[len(this.stack)-1]
}


func (this *MinStack) GetMin() int {
    var min int
    for i, v := range this.stack{
        if i == 0 {
            min = v
        }

        if v < min {
            min = v
        }
    }
    return min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */