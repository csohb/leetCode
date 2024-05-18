func twoSum(numbers []int, target int) []int {
    var answer []int

    left, right := 0, len(numbers) - 1
    for left < right {
        l := numbers[left]
        r := numbers[right]

        fmt.Println("l:",l)
        fmt.Println("r:",r)
        if l+r == target {
            answer = []int{left+1, right+1}
            break
        }

        if l+r > target {
            right--
        } else {
            left++
        }
    }

    return answer
}