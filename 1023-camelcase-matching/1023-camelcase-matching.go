func camelMatch(queries []string, pattern string) []bool {
    res := []bool{}
    for _, query := range queries {
        var pIndex int
        ok := true
        for _, s := range query {
            if pIndex < len(pattern) {
                if s == rune(pattern[pIndex]) {
                    pIndex++
                    continue
                }
            } 
            if unicode.IsUpper(s) {
                ok = false
                break
            }
        }
        res = append(res, ok && len(pattern) == pIndex)
    }
    return res
}


func camelToSnake(s string) string {
    res := []rune{}
    for i, v := range s {
        if unicode.IsUpper(v) {
            if i > 0 {
                res = append(res, '_')
            }
            res = append(res, unicode.ToLower(v))
        } else {
            res = append(res, v)
        }
    }

    return string(res)
}