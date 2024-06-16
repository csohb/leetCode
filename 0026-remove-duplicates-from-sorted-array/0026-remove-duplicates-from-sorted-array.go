func removeDuplicates(nums []int) int {
    var j int
    numMap := make(map[int]int, 0)
    for _, v := range nums {
        if _, ok := numMap[v]; !ok {
            numMap[v] = 1
            nums[j] = v
            j++
        }
    }
    return j
}