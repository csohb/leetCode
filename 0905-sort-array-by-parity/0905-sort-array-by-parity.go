func sortArrayByParity(nums []int) []int {
    // check all elments in nums arr -> O(n)
        // check if it's odd or even 
        // num % 2 == 0 -> even / odd 
    // create 2 array 
    // odd arr : m 
    // even arr : n - m 
    // not exceed O(n) 
    // even = append(even, odd)


    evenArr := make([]int, 0)
    oddArr := make([]int, 0)

    for _, num := range nums {
        if num % 2 == 0 {
            // even
            evenArr = append(evenArr, num)
        } else {
            oddArr = append(oddArr, num)
        }
    }

    //evenArr = append(evenArr, oddArr)

    for _, oddNum := range oddArr {
        evenArr = append(evenArr, oddNum)
    }


    return evenArr

}