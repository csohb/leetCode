func groupAnagrams(strs []string) [][]string {
    // map check if all char is included
    m := make(map[string][]string)

    for _, str := range strs {
        // split all char
        chars := strings.Split(str, "")
        // sort it asc
        sort.Strings(chars)
        // join to make key char ex) aet
        key := strings.Join(chars, "")
        m[key] = append(m[key], str)
    }

    res := make([][]string, 0, len(m))
    for _, v := range m {
        res = append(res, v)
    }

    return res
}