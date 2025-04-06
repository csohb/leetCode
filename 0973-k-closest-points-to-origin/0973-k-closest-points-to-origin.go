func kClosest(points [][]int, k int) [][]int {
    res := make([][]int, 0)

    // i < k --> append to List (sort), append to res

    // array
    // dList := make([][]int, 0)

    // for _, point := range points {
    //     distance := (point[0] * point[0]) + (point[1] * point[1])
    //     dList = append(dList, []int{distance, point[0], point[1]})        
    // }

    // sort.Slice(dList, func(i, j int) bool{
    //     return dList[i][0] < dList[j][0]
    // })

    // for i:=k-1; i>=0; i-- {
    //     x := dList[i][1]
    //     y := dList[i][2]
    //     res = append(res, []int{x,y})
    // }

    minHeap := &MinHeap{}
    heap.Init(minHeap)

    for _, point := range points {
        x, y := point[0], point[1]
        distance := (x * x) + (y * y)
        heap.Push(minHeap, Point{distance:distance, x:x, y:y})
    }

    for k > 0 {
        point := heap.Pop(minHeap).(Point)
        res = append(res, []int{point.x, point.y})
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

func (h MinHeap) Len() int {
    return len(h)
}

func (h MinHeap) Less(i, j int) bool {
    return h[i].distance < h[j].distance
}

func (h MinHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
    *h = append(*h, x.(Point))
}

func (h *MinHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0:n-1]
    return x
}