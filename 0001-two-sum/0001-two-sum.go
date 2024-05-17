func twoSum(nums []int, target int) []int {
    // two numbers 
    var result []int
    targetMap := make(map[int]int)

    for i, v := range nums {
        diff := target - v
        if vv, ok := targetMap[diff]; ok {
            result = []int{i, vv}
        }
        targetMap[v] = i
    }

    return result
}