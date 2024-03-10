func minimumBoxes(apple []int, capacity []int) int {
    // greedy sum all the apples and also capacities
    // apples are 
    var result int

    var appleSum int
    for _, v := range apple {
        appleSum += v
    }
    
    sort.Slice(capacity, func(i, j int) bool {
    return capacity[i] > capacity[j]
    })

    for appleSum > 0 {
        appleSum -= capacity[result]
        result++
    }
     
       
    return result
}
