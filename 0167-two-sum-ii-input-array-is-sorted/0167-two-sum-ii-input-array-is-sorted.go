func twoSum(numbers []int, target int) []int {
    // two sum ... 
    end := len(numbers) - 1
    answer := make([]int, 0, 2)

    for i:=0; i < len(numbers); i++ {
        for j:= end; j >= 0; j-- {
            if numbers[i] + numbers[j] == target {
                answer = append(answer, i+1)
                answer = append(answer, j+1)
                return answer
            }
        }    
    } 

    return answer
}