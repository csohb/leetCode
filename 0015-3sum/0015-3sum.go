func threeSum(nums []int) [][]int {
    /*
        fix x index
        make sum with left, right
        if 0, then add to answer
        if not, move left or right
    */
    var answer [][]int
    sort.Ints(nums)
    for i:=0; i < len(nums); i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        x := nums[i]
        left, right := i+1, len(nums)-1
        for left < right {
            y := nums[left]
            z := nums[right]
            if x + y + z == 0 {
                answer = append(answer, []int{x, y, z})
                left++
                for nums[left] == nums[left-1] && left < right {
                    left++
                }
                continue
            }

            if x + y + z < 0 {
                left++
            } else {
                right--
            }
        }
    }
    return answer
}