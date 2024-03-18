func removeDuplicates(nums []int) int {
    // more than twice duplicated then delete it
    
    var curNum int
    var count int
    var j int

    curNum = nums[0]
    count++
    j++

    for i, v := range nums {
        if i == 0 {
            continue
        }
        
        if curNum == v {
            if count < 2 {
                nums[j] = v
                count++
                j++
                continue
            } else {
                continue
            }
        } else {
            if count > 0 {
                count = 0
            }
            curNum = v
            nums[j] = v
            count++
            j++
        }
    }

    return j
}