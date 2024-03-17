func countSubstrings(s string, c byte) int64 {
    var count int64
    
    // for i:=0; i < len(s); i++ {
    //     if s[i] == c {
    //         for j := i + 1; j <= len(s); j++ {
    //             if s[j-1] == c {
    //                 count++
    //             }
    //         }
    //     }
    // }
    var answer int64
    for i:=0; i < len(s); i++ {
        if s[i] == c {
            count++
            answer+=count
        }
    }
    
    return answer
}