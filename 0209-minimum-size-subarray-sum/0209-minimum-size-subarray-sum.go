func minSubArrayLen(target int, nums []int) int {
    var minLen = len(nums) + 1 // initialize minLen to a value greater than the length of nums
    var left, sum int

    for right, num := range nums {
        sum += num
        // Shrink the window as much as possible while the sum is greater than or equal to target
        for sum >= target {
            minLen = min(minLen, right-left+1) // Update minLen
            sum -= nums[left] // Shrink the window from the left
            left++
        }
    }

    // If minLen is still greater than the length of nums, it means no subarray was found
    if minLen > len(nums) {
        return 0
    }

    return minLen
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
