func removeStars(s string) string {
    sByte := []byte(s)

    newStr := strStack{}
    for _, v := range sByte {
        if v != '*' {
            // push v to stack 
            newStr.Push(v)
            //fmt.Println("newStr:", newStr)
            //newStr += string(v)
        } else {
            // pop a character from stack
            newStr.Pop()
            //newStr = newStr[:len(newStr) - 1]
        }  
    }

    return string(newStr.arr)
}

type strStack struct {
    arr []byte
}

func (s *strStack) Push(char byte) {
    s.arr = append(s.arr, char)
} 

func (s *strStack) Pop() {
    s.arr = s.arr[:len(s.arr)-1]
}