func twoSum(nums []int, target int) []int {
    res := make([]int, 0)
    for i, v := range nums {
       
        sum := v

        for j:=i+1; j<len(nums); j++ {
            if sum + nums[j] == target {
                res = append(res, i)
                res = append(res, j)
            }
        }
    }

    return res
}