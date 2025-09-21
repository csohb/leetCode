func longestPalindrome(s string) string {
    res := ""
    resLen := 0

    for i, _ := range s {
        // odd case
        left, right := i, i 
        for left >= 0 && right < len(s) && s[left] == s[right] {
            if right - left + 1 > resLen {
                res = s[left:right+1]
                resLen = right - left + 1
            }
            left--
            right++
        }

        // even case
        left, right = i, i+1 
        for left >= 0 && right < len(s) && s[left] == s[right] {
            if right - left + 1 > resLen {
                res = s[left:right+1]
                resLen = right - left + 1
            }
            left--
            right++
        }
    }
    return res
}

