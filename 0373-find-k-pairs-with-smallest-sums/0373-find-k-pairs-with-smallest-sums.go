func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
    var result [][]int
    if len(nums1) == 0 || len(nums2) == 0 || k <= 0 {
        return result
    }
    pq := &PriorityQueue{}
    heap.Init(pq)
    for i := 0; i < len(nums1) && i < k; i++ {
        for j := 0; j < len(nums2) && j < k; j++ {
            sum := nums1[i] + nums2[j]
            pair := Pair{
                pairSum: sum,
                nums:  []int{nums1[i], nums2[j]},
            }
            if pq.Len() < k {
                heap.Push(pq, pair)
            } else if sum < (*pq)[0].pairSum {
                heap.Pop(pq)
                heap.Push(pq, pair)
            } else {
                break
            }
        }
    }
    for pq.Len() > 0 {
        pair := heap.Pop(pq).(Pair)
        result = append(result, pair.nums)
    }
    return result
}

type Pair struct {
    pairSum int
    nums []int
}

type PriorityQueue []Pair

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
    return pq[i].pairSum > pq[j].pairSum
}

func (pq PriorityQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
    item := x.(Pair)
    *pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
    old := *pq
    n := len(old)
    item := old[n-1]
    *pq = old[: n-1]
    return item
}