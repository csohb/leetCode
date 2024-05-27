
func evalRPN(tokens []string) int {
    numStack := &NumStack{}
    for _, v := range tokens {
        num, err := strconv.Atoi(v)
        if err != nil {
            // string
            num1 := numStack.pop()
            num2:= numStack.pop()

            fmt.Println("num1:",num1,"num2:",num2)
            var answer int
            switch v {
            case "+":
                answer = num2 + num1
            case "*":
                answer = num2 * num1
            case "-":
                answer = num2 - num1
            case "/":
                answer = num2 / num1
            }
            numStack.push(answer)
        } else {
            // num
            numStack.push(num)
        }
    }

    return numStack.stack[0]
}

type NumStack struct {
    stack []int
}

func (n *NumStack) push(num int) {
    n.stack = append(n.stack, num)
}

func (n *NumStack) pop() int {
    top := n.stack[len(n.stack)-1]
    n.stack = n.stack[:len(n.stack)-1]
    return top
}

func max(a, b int) (bigger, smaller int) {
    if a > b {
        return a, b
    } 
    return b, a
}