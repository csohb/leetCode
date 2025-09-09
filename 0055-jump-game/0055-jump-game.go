func canJump(nums []int) bool {
    // how far we can reach : farhest
    // update max
    // 0  
    var max int
    for index:=0; index<len(nums); index++ {
        if index > max {
            return false
        }
        // index :0, nums[index]: 2
        if index + nums[index] > max {
            max = index + nums[index]
        }

        if max >= len(nums) - 1 {
            return true
        }
    }

    return true
}