func isValid(s string) bool {
    // stack 
    // open brackets are pushed to stack
    // whenever s meets closing brackets then compare with last input data
    // stack[len(stack)] -> if it's pair then continue otherwise false

    if len(s) % 2 == 1 {
        return false
    }

    stack := []rune{}

    for _, ss := range s {
        if ss == '(' || ss == '{' || ss == '[' {
            stack = append(stack, ss)
        } else {
            if len(stack) <= 0 {
                return false
            }
            last := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            switch ss {
            case ')': 
                if last != '(' {
                    return false
                }
            case '}':
                if last != '{' {
                    return false
                }
            case ']':
                if last != '[' {
                    return false
                }
            }
        }
    }

    if len(stack) > 0 {
        return false
    }
    return true
}