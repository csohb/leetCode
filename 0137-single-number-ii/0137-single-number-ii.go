func singleNumber(nums []int) int {
    var res int

    checkMap := map[int]int{}

    for _, v := range nums {
        checkMap[v]++
    }

    for k, v := range checkMap {
        if v == 1 {
            res = k
            break
        }
    }

    return res
}