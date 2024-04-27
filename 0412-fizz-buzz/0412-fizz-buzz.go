func fizzBuzz(n int) []string {
    // fizzMap := make(map[int]string, 3)
    // fizzMap[3] = "Fizz"
    // fizzMap[5] = "Buzz"
    // fizzMap[15] = "FizzBuzz"

    var answer []string
    for i:=1; i<=n; i++ {
        if i % 15 == 0 {
            answer = append(answer, "FizzBuzz")
            continue
        } else if i % 3 == 0 {
            answer = append(answer, "Fizz")
            continue
        } else if i % 5 == 0 {
            answer = append(answer, "Buzz")
            continue
        }

        answer = append(answer, strconv.Itoa(i))
    } 

    return answer
}