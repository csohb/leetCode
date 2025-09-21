func canJump(nums []int) bool {
    // save max reach
    var max int

    for i:=0; i<len(nums); i++ {
        if i > max {
            return false
        }

        if i + nums[i] > max {
            max = i + nums[i]
        }

        if max >= len(nums) - 1{
            return true
        }
    }

    return true
}