func twoSum(numbers []int, target int) []int {
    // two sum ... 
    start := 0
    end := len(numbers) - 1
    answer := make([]int, 0, 2)

    for end > start {
        if numbers[start] + numbers[end] > target {
            end--
        } else if numbers[start] + numbers[end] < target{
            start++
        } else {
            answer = []int{start+1, end+1}
            return answer
        }
    }
    return answer
}