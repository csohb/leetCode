func twoSum(numbers []int, target int) []int {
    // two sum ... 

    var answer []int
    for i, v := range numbers {
        for j, vv := range numbers {
            if i == j {
                continue
            }

            if v + vv == target {
                //fmt.Println("v:", v, ",vv:",vv)
                answer = append(answer, i+1)
                answer = append(answer, j+1)
                return answer
            }
        }
    } 

    return answer
}