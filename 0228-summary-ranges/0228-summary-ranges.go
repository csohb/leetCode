
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
            result = append(result, makeResultStr(currRange))
            break
        }
        numsIndex++
        start++
    } else {
        if len(currRange) < 1 {
                start++
                continue
        }
        result = append(result, makeResultStr(currRange))
        currRange = []int{}
        start = nums[numsIndex]
        continue
        // if len(currRange) == 1 {
        //         result = append(result, fmt.Sprintf("%d",currRange[0]))
        //         currRange = []int{}
        //         start = nums[numsIndex]
        //         continue
        //     }
            
        //     result = append(result, fmt.Sprintf("%d->%d",currRange[0], currRange[len(currRange)-1]))
        //     currRange = []int{}
        //     start = nums[numsIndex]
        //     continue
    }
   }

   return result 
}

func makeResultStr(currRange []int) string {
    if len(currRange) == 1 {
        return fmt.Sprintf("%d",currRange[0])
    } else {
        return fmt.Sprintf("%d->%d",currRange[0], currRange[len(currRange)-1])
    }
}