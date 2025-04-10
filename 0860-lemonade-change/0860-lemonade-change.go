func lemonadeChange(bills []int) bool {
    change := make(map[int]int)
    for i:=0; i < len(bills); i++ {
        switch bills[i] {
        case 5:
            change[5]++
        case 10:
            change[10]++
            if change[5] >= 1 {
                change[5]--
            } else {
                return false
            }
        case 20:
            fmt.Println("change:", change)
            if change[5] >= 1 && change[10] >= 1 {
                change[10]--
                change[5]-=1
            } else if change[5] >= 3 {
                change[5]-=3
            } else {
                return false
            }
        }
        
    }
    return true
}