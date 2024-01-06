package main

import (
    "fmt"
    "sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
    nums1 = nums1[:m]
    nums2 = nums2[:n]
    fmt.Println("nums2:", nums2)
    nums3 := append(nums1, nums2...)
    sort.Ints(nums3)
    nums1 = nums3
}

func main() {
    nums1 := []int{1, 2, 3, 0, 0, 0}
    m := 3
    nums2 := []int{2, 5, 6}
    n := 3
    merge(nums1, m, nums2, n)
    fmt.Println(nums1)
}
