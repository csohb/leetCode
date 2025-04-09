func minPairSum(nums []int) int {
    // too make minimun max pair --> biggest + smallest
    sort.Slice(nums, func(i, j int) bool {
        return nums[i] > nums[j]
    }) 

    length := len(nums) / 2
    start := 0
    end := len(nums)-1

    h := &MaxHeap{}
    heap.Init(h)
    for length > 0 {
        big := nums[start]
        small := nums[end]
        x := MaxPair{sum:big+small}
        heap.Push(h, x)
        start++
        end--
        length--
    }

    x := heap.Pop(h).(MaxPair)

    return x.sum
    
}

type MaxPair struct {
    sum int
}

type MaxHeap []MaxPair

func(h MaxHeap) Len() int {
    return len(h)
}

func(h MaxHeap) Less(i, j int) bool {
    return h[i].sum > h[j].sum
}

func(h MaxHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func(h *MaxHeap) Push (x interface{}) {
    *h = append(*h, x.(MaxPair))
}

func(h *MaxHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}