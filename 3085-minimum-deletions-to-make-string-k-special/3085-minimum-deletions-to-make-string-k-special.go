func minimumDeletions(word string, k int) int {
    // 가장 많이 나온 알파벳의 횟수 - 가장 적게 나온 알파벳의 횟수 <= k 
    // 이걸 만드는 가장 작은 int 

    count := 99999

    alphaMap := make(map[rune]int, len(word))
    
    for _, v := range word {
        if _, ok := alphaMap[v]; !ok {
            alphaMap[v] = 1
        } else {
            alphaMap[v]++
        }
    }
    
    var countSlice []int
    
    for _, v := range alphaMap {
        countSlice = append(countSlice, v)
    }

    sort.Slice(countSlice, func(i, j int) bool{
        return countSlice[i] > countSlice[j]
    })

    for i:= 0; i < len(countSlice); i++ {
        curCount := countSlice[i]
        delCount := 0
        for j:= 0; j < len(countSlice); j++ {
            if countSlice[j] > curCount + k {
                delCount += countSlice[j] - (curCount + k)
            } else if (countSlice[j] < curCount) {
                delCount += countSlice[j]
            }
        }
         if delCount < count {
            count = delCount
        }
    }

    return count
}