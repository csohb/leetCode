func reverseString(s []byte)  {
    half := len(s) / 2 
    start := 0
    end := 1 
    for half > 0 {
        s[start], s[len(s)-end] = s[len(s)-end], s[start]
        start++
        end++
        half--
    }
}