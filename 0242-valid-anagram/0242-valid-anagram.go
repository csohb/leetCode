func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    sMap := make(map[byte]int)
    tMap := make(map[byte]int)

    for i:=0; i < len(s); i++ {
        sMap[s[i]]++
        tMap[t[i]]++
    }

    for k, v := range sMap {
        if tMap[k] != v {
            return false
        }
    }

    return true
}