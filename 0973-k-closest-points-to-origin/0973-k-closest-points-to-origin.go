func kClosest(points [][]int, k int) [][]int {
    // *Minimum or Closest -> Min Heap
    res := make([][]int, 0)

    mHeap := &MinHeap{}
    heap.Init(mHeap)

    for _, point := range points {
        x := point[0]
        y := point[1]
        distance := (x*x) + (y*y)

        heap.Push(mHeap, Point{distance, x, y})
    }


    for k > 0 {
        point := heap.Pop(mHeap).(Point)
        res = append(res, []int{point.x,point.y})
        k--
    }

    return res
}

type Point struct {
    distance int
    x int
    y int
}

type MinHeap []Point

func(h MinHeap) Len() int {
    return len(h)
}

func(h MinHeap) Less(i, j int) bool {
    return h[i].distance < h[j].distance
}

func(h MinHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func(h *MinHeap) Push(x interface{}) {
    *h = append(*h, x.(Point))
}

func(h *MinHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0:n-1]
    return x
}
