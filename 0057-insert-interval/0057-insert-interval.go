func insert(intervals [][]int, newInterval []int) [][]int {
    var answer [][]int
    if len(intervals) == 0 && len(newInterval) == 0 {
        return answer
    }

    if len(newInterval) > 0 {
        intervals = append(intervals, newInterval)
    } 

    if len(intervals) == 1 {
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
    // 1.last < 2.start
    // 
    for i, v := range intervals {
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