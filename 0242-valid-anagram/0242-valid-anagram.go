func isAnagram(s string, t string) bool {
    // store alphabet in hashMap and compare if it's same or not
    if len(s) != len(t) {
        return false
    }

    // 2 hashMap
    sMap := make(map[byte]int)
    tMap := make(map[byte]int)
    for i:=0; i<len(s); i++ {
        if _, ok := sMap[byte(s[i])]; ok {
            sMap[byte(s[i])]++
        } else {
            sMap[byte(s[i])] = 1
        }
    }

    for i:=0; i<len(t); i++ {
        if _, ok := tMap[byte(t[i])]; ok {
            tMap[byte(t[i])]++
        } else {
            tMap[byte(t[i])] = 1
        }
    }

    for k, v := range sMap {
        if v != tMap[k] {
            return false
        }
    }

    return true
}