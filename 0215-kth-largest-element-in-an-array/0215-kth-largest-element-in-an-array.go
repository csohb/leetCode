type MinHeap []int

func(m MinHeap) Less(i, j int) bool {
    return m[i] > m[j]
}

func(m MinHeap) Swap(i, j int) {
    m[i], m[j] = m[j], m[i]
}

func(m MinHeap) Len() int {
    return len(m)
}

func(m *MinHeap) Push(v any) {
    *m = append(*m, v.(int))
}

func(m *MinHeap) Pop() any {
    old := *m
    n := len(old)
    x := old[n-1]
    *m = old[0:n-1]
    return x
}

func findKthLargest(nums []int, k int) int {
    h := &MinHeap{}
    for i:=0; i< len(nums); i++ {
        heap.Push(h, nums[i])
    }

    var result int
    for k > 0 {
        result = heap.Pop(h).(int)
        k--
    }

    return result
}