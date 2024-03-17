func merge(nums1 []int, m int, nums2 []int, n int)  {
    var index int
    for i, _ := range nums1 {
        if i >= m {
            nums1[i] = nums2[index]
            index++
        }
    }
    sort.Slice(nums1, func(i, j int) bool{
        return nums1[i] < nums1[j]
    })
}