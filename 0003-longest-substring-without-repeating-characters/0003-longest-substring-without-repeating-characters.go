func lengthOfLongestSubstring(s string) int {
    if len(s) < 1 {
        return 0
    }

    max := 1
    
    for i, _ := range s {
        charMap := make(map[byte]bool)
        charMap[s[i]] = true
        //fmt.Println("s[i]:", string(s[i]))
        cnt := 1
        for ii := i + 1; ii < len(s); ii++ {
            if charMap[s[ii]] == false {
                //fmt.Println("s[ii]:", string(s[ii]))
                charMap[s[ii]] = true
                cnt++
                //fmt.Println("cnt:",cnt)
                if cnt > max {
                    max = cnt
                }
            } else {
                break
            }
        }
    }   

    return max
}