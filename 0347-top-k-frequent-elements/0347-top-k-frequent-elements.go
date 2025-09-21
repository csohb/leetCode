func topKFrequent(nums []int, k int) []int {
    // min or max -> using data structure "Heap"
    // heap push, pop will be O(1) -> constance time
    // even though making a heap from the list will be O(n)    
    freq := make(map[int]int)
    for _, v := range nums {
        freq[v]++
    }

    kHeap := &MaxHeap{}
    heap.Init(kHeap)

    for k, v := range freq {
        heap.Push(kHeap, Pair{k, v})
    }

    res := []int{}
    for k > 0 {
        res = append(res, heap.Pop(kHeap).(Pair).key)
        k--
    }

    return res
}

type Pair struct {
    key int
    freq int
}

type MaxHeap []Pair

func(h MaxHeap) Len() int {
    return len(h)
}

func(h MaxHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func(h MaxHeap) Less(i, j int) bool {
    return h[i].freq > h[j].freq
}

func(h *MaxHeap) Push(x interface{}) {
    *h = append(*h, x.(Pair))
}

func(h *MaxHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}