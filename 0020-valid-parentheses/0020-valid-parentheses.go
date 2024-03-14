func isValid(s string) bool {
    if len(s) % 2 != 0 || len(s) == 0{
        return false
    }

    bracketMap := map[rune]rune {
        '}': '{',
		']': '[',
		')': '(',
    }

    bracket := NewBracketStack(len(s))

    for _, char := range s {
        opening, ok := bracketMap[char] 
        if !ok {
            bracket.Push(char)
            continue
        }

        // if closing
        if len(bracket.stack) == 0 {
            return false
        }

        if bracket.stack[len(bracket.stack)-1] != opening {
            return false
        }

        // pop opening bracket
        bracket.Pop()
    }

    return len(bracket.stack) == 0
}

type bracketStack struct {
    stack []rune
}

func NewBracketStack(sLen int) *bracketStack {
    return &bracketStack{
        stack : make([]rune, 0, sLen),
    }
}

func(b *bracketStack) Push(sChar rune) {
    b.stack = append(b.stack, sChar)
}

func(b *bracketStack) Pop() {
    b.stack = b.stack[:len(b.stack)-1]
}

