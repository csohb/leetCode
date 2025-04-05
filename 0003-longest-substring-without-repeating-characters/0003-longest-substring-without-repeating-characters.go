func lengthOfLongestSubstring(s string) int {
    // first answer
    /*if len(s) < 1 {
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

    return max*/

    // sliding window
    // charSet := make(map[byte]bool)
    // l := 0
    // res := 0

    // for r := 0; r < len(s); r++ {
    //     for charSet[s[r]] {
    //         delete(charSet, s[l])
    //         l++
    //     }
    //     charSet[s[r]] = true
    //     res = max(res, r-l+1)
    // }
    
    // return res


    l := 0
    res := 0

    charMap := make(map[byte]bool)
    
    for r:=0; r < len(s); r++ {
        for charMap[s[r]] {
            delete(charMap,s[l])
            l++
        }
        charMap[s[r]] = true
        res = max(res, r-l+1)
    }

    return res
}