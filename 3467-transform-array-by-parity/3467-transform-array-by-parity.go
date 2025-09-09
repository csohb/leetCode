func transformArray(nums []int) []int {
    // 1. brute force: check all the elements -> replace 0, 1
    // time complexity -> O(n), 
    // space complexity -> O(1)

    // 1, 2 steps
    for i, num := range nums {
        if num % 2 == 0 {
            nums[i] = 0
        } else {
            nums[i] = 1
        }
    }

    // step 3
    // sort.Slice -> O(n)
    sort.Slice(nums, func(i, j int) bool{
        return nums[i] < nums[j]
    })

    return nums

    // 2. count how many even / odd numbers -> create new arr with cnt 
    // time complexity -> O(n)
    // space complexity -> O(n)
}