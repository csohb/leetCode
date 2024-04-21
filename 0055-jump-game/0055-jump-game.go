func canJump(nums []int) bool {
    if len(nums) == 0 {
        return false
    }

    var i int
    var possible int
    for ; i < len(nums); i++ {
        // max
        if i + nums[i] > possible {
            possible = i + nums[i]
        }

        if nums[i] == 0 && i < len(nums) - 1 && i == possible {
            return false
        }
    }

    
    return true
}
