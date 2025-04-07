func removeDuplicates(nums []int) int {
    if len(nums) <= 0 {
        return 0
    }

    dup:=0
    newNum := []int{}
    curr := nums[0]
    newNum = append(newNum, curr)

    // count duplicate 
    for i:=1; i < len(nums); i++ {
        if nums[i] == curr {
            dup++
            continue
        }
        newNum = append(newNum, nums[i])
        curr = nums[i]
    }



    for i:=0; i<len(newNum); i++ {
        nums[i] = newNum[i]
    }

    res := len(newNum)
    return res
}