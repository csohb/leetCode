func merge(intervals [][]int) [][]int {
    var answer [][]int
    // start, last 
    // 2.start <= 1.last -> 1.last 2.last 더 큰걸로 last
    // 2.start > l.last // new list 

    if len(intervals) == 0 {
        return answer
    } else if len(intervals) == 1 {
        return intervals
    }

    sort.Slice(intervals, func(i, j int) bool {
        if intervals[i][0] == intervals[j][0] {
            return intervals[i][1] < intervals[j][1]
        }
        return intervals[i][0] < intervals[j][0]
    })

    var start int
    var last int
    var curr []int
    for i, v := range intervals {
        // fmt.Println("start:",start)
        // fmt.Println("last:",last)
        // fmt.Println("curr:",curr)
        // fmt.Println("v:",v)
        if i == 0 {
            start = v[0]
            last = v[1]
            continue
        }

        if v[0] <= last {
            if last < v[1] {
                last = v[1]
            }
            if i == len(intervals) -1 {
                curr = append(curr, start, last)
                answer = append(answer, curr)
                break   
            }
        } else if v[0] > last {
            curr = append(curr, start, last)
            answer = append(answer, curr)
            curr = nil
            start = v[0]
            last = v[1]
            if i == len(intervals) -1 {
                curr = append(curr, start, last)
                answer = append(answer, curr)
                break   
            }
            continue
        }
    }

    return answer
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}