func twoSum(nums []int, target int) []int {
    res := make([]int, 0)
    for i, v := range nums {
       
        sum := v

        for j:=i+1; j<len(nums); j++ {
            if sum + nums[j] == target {
                fmt.Println("i:",nums[i])
                fmt.Println("j:",nums[j])
                res = append(res, i)
                res = append(res, j)
            }
        }
    }

    return res
}