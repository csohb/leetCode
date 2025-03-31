func longestConsecutive(nums []int) int {
    m := map[int]bool{}
    for _, num := range nums {
        m[num] = true
    }
    res := 0
    for num, _ := range m {
        if !m[num-1] {
            start := num
            cnt := 0
            for m[start] {
                cnt++
                start++
            }
            if cnt > res {
                res = cnt
            }
        }
    }

    return res
}