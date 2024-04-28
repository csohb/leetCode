func searchInsert(nums []int, target int) int {
    n := len(nums)
    low, high := 0, n - 1
    for low < high {
        mid := (low + high) / 2
        // fmt.Println("mid:",mid)
        // fmt.Println("low:",low)
        // fmt.Println("high:",high)
        if nums[mid] < target {
            low = mid + 1
        } else {
            high = mid
        }
    }

    if nums[low] < target {
        return low + 1
    } else {
        return low
    }
}