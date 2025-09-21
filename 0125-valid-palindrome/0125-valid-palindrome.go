func isPalindrome(s string) bool {
    // convert all char -> lower alphabet if can
    // make new string 
    // check new string if it's palindrome or not

    // check palindrome
    // two pointers: start, end 
    // if newS[start] == newS[end] -> keep going 
    // if not false

    // update all char in s -> valid lower char

    newS := ""
    for _, v := range s {
        if unicode.IsLetter(v) || unicode.IsDigit(v) {
            newS += string(unicode.ToLower(v))
        }
    }

    fmt.Println("newS:", newS)

    // expect new String only with lower characters

    start := 0
    end := len(newS) - 1
    mid := len(newS) / 2

    for mid > 0 {
        if newS[start] == newS[end] {
            start++
            end--
            
            mid--
            continue
        } else {
            return false
        }
    }

    return true
}