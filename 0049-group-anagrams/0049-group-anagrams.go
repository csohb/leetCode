func groupAnagrams(strs []string) [][]string {
    // for loop n times n = len(strs) : time complexity O(n)
    // group with same annagram
    // new str -> need to find if it's exists or not (group)
        // hashMap to find the key : time complexity O(1)
    // if yes then add in the arr 
    // if no then create new arr and add in result

    res := make([][]string, 0, len(strs))
    hMap := make(map[string][]string)

    for _, str := range strs {
        sortArr := make([]byte, 0)
        for _, alphabet := range str {
            // byte
            sortArr = append(sortArr, byte(alphabet))
        }
        sort.Slice(sortArr, func(i, j int) bool{
            return sortArr[i] < sortArr[j]
        })
        sortedWord := string(sortArr)
        
        if _, ok := hMap[sortedWord]; ok {
            hMap[sortedWord] = append(hMap[sortedWord], str)
        } else {
            hMap[sortedWord] = append([]string{}, str)
        }
    }

    for _, v := range hMap {
        res = append(res, v)
    }
    return res
}