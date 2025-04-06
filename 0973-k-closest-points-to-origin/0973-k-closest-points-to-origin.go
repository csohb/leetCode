func kClosest(points [][]int, k int) [][]int {
    res := make([][]int, 0)

    // i < k --> append to List (sort), append to res

    dList := make([][]int, 0)

    for _, point := range points {
        distance := (point[0] * point[0]) + (point[1] * point[1])
        dList = append(dList, []int{distance, point[0], point[1]})        
    }

    sort.Slice(dList, func(i, j int) bool{
        return dList[i][0] < dList[j][0]
    })

    for i:=k-1; i>=0; i-- {
        x := dList[i][1]
        y := dList[i][2]
        res = append(res, []int{x,y})
    }

    return res
}
