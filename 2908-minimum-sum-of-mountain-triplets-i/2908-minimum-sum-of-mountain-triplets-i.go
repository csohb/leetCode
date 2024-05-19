func minimumSum(nums []int) int {
    var sum int
    sum = math.MaxInt
    for i, _ := range nums {
        iNum := nums[i]
        for j := i + 1; j < len(nums) -1; j++ {
            jNum := nums[j]
            for k := j + 1; k < len(nums); k++ {
                kNum := nums[k]
                if iNum < jNum && kNum < jNum {
                    sum = min(sum, iNum+jNum+kNum)
                }
            }
        }

    }

    if sum == math.MaxInt {
        sum = -1
    }
    return sum
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}