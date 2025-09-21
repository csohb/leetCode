func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }

    // 0 ~ 9 --> true
    if x < 10 {
        return true
    }

    arr := []int{}
    //cnt := 0
    for x > 0 {
        num := x % 10
        arr = append(arr, num)
        
        x = x / 10
        //fmt.Println("x:",x)
    }

    mid := len(arr) / 2
    start := 0
    end := len(arr) - 1

    for i:=0; i<mid; i++ {
        if arr[start] == arr[end] {
            // palindrome
            start++
            end--
            continue
        } else {
            return false
        }
    }

    //fmt.Println("arr:",arr)

    
    // 10 ~ 99 --> only 11, 22, ... 99 is true
    // 101 ~ 999 --> 

    return true
    
}