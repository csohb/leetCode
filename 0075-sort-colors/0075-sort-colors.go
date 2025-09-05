func sortColors(nums []int)  {
    // min value should come first
    colorMap := map[int]int{}
    for _, num := range nums {
        colorMap[num]++
    }

    for i := range nums {
        if colorMap[0] > 0 {
            nums[i] = 0
            colorMap[0]--
        } else if colorMap[1] > 0 {
            nums[i] = 1
            colorMap[1]--
        } else {
            nums[i] = 2
            colorMap[2]--
        }
    }
}