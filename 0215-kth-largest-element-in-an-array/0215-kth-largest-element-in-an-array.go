func findKthLargest(nums []int, k int) int {
    // max or min -> heap 
    // when input, it will be sorted 
    // time complexity will be O(1)
    numHeap := &MaxHeap{}
    heap.Init(numHeap)

    for _, v := range nums {
        heap.Push(numHeap, v)
    }   

    var res int
    for k > 0 {
        res = heap.Pop(numHeap).(int)
        k--
    }

    return res
}

type MaxHeap []int

func(h MaxHeap) Len() int {
    return len(h)
}

func(h MaxHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func(h MaxHeap) Less(i, j int) bool {
    return h[i] > h[j]
}

func(h *MaxHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

func(h *MaxHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}