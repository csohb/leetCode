func twoSum(nums []int, target int) []int {
    // only one solution -> when solution founded return
    output := []int{}

    for i, _ := range nums {
        for ii:= i + 1; ii < len(nums); ii++ {
            if nums[i] + nums[ii] == target {
                output = append(output, i, ii)
                return output
            }
        }
    }

    return output
}