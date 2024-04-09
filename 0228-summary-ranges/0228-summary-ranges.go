
func summaryRanges(nums []int) []string {
   var result []string
   if len(nums) < 1 {
        return result 
    }
   numsMap := make(map[int]bool,0)
   for _, v := range nums {
        numsMap[v] = true
   }
   var numsIndex int
   var currRange []int

   start := nums[0]

   for start <= nums[len(nums)-1] {
    if numsMap[start] {
        currRange = append(currRange, start)
        // 마지막 숫자
        if start == nums[len(nums) -1] {
            if len(currRange) == 1 {
                result = append(result, fmt.Sprintf("%d",currRange[0]))
            } else {
                result = append(result, fmt.Sprintf("%d->%d",currRange[0], currRange[len(currRange)-1]))
            }
            break
        }
        numsIndex++
        start++
    } else {
        if len(currRange) < 1 {
                start++
                continue
        }
        if len(currRange) == 1 {
                fmt.Println("currRange:",currRange)
                result = append(result, fmt.Sprintf("%d",currRange[0]))
                currRange = []int{}
                start = nums[numsIndex]
                continue
            }
            
            result = append(result, fmt.Sprintf("%d->%d",currRange[0], currRange[len(currRange)-1]))
            currRange = []int{}
            start = nums[numsIndex]
            fmt.Println("start:",start)
            continue
    }
   }

//    for i:= nums[0]; i<=nums[len(nums)-1]+1; i++ {
//         fmt.Println("start i:",i)
//         if numsMap[i] {
//             currRange = append(currRange, i)
//             numsIndex++
//             fmt.Println("currRange:",currRange)
//             continue
//         } else {
//             if len(currRange) < 1 {
//                 continue
//             }
//             if len(currRange) == 1 {
//                 fmt.Println("currRange:",currRange)
//                 result = append(result, fmt.Sprintf("%d",currRange[0]))
//                 currRange = []int{}
//                 i = nums[numsIndex]
//                 continue
//             }
            
//             result = append(result, fmt.Sprintf("%d->%d",currRange[0], currRange[len(currRange)-1]))
//             currRange = []int{}
//             i = nums[numsIndex]
//             fmt.Println("i:",i)
//             continue
//         }
//    }
   return result 
}