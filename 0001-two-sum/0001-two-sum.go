func twoSum(nums []int, target int) []int {
    // two numbers 
    var result []int 

    // target 에다가 하나씩 빼면서 0이 되는지 체크 
    // map ... index, value 추가 

    sumMap := make(map[int]int)
    for i, v := range nums {
        sumMap[i] = v
    }


    for k, v := range sumMap {
        for k2, v2 := range sumMap {
            if k2 == k {
                continue
            }
            if v + v2 == target {
                result = []int{k, k2}
            }
        }
    }

    return result
}