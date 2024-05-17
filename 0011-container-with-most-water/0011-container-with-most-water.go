func maxArea(height []int) int {
   var answer int
//    for i := 0; i < len(height); i++ {
//     left := height[i]
//     for j:= i+1; j < len(height); j++ {
//         right := height[j]
//         minVal := min(left, right)
//         square := minVal * (j - i)
//         answer = max(square, answer)
//     }
//    }
    left, right := 0, len(height)-1
    // until they cross over
    for left < right {
        l := height[left]
        r := height[right]
        minVal := min(l, r)
        square := minVal * (right - left)
        answer = max(square, answer)
        if l < r {
            left++
        } else {
            right--
        }
    }

   return answer 
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}