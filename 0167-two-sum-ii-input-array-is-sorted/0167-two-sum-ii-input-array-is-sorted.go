func twoSum(numbers []int, target int) []int {
    // two sum ... 

    answer := make([]int, 0, 2)
    for i, v := range numbers {
        for j, vv := range numbers {
            if i == j {
                continue
            }

            if v + vv == target {
                answer = append(answer, i+1)
                answer = append(answer, j+1)
                return answer
            }
        }
    } 

    return answer
}