func subarraySum(nums []int, k int) int {
    cnt := 0
    for i, _ := range nums {
        sum := 0
        for j:=i; j<len(nums); j++ {
            sum+= nums[j]
            if sum == k {
                cnt++
            }
        }
    }
    return cnt
}