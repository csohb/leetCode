func multiply(num1 string, num2 string) string {
    // if num1 == "0" || num2 == "0" {
    //     return "0"
    // }

    // resultArr := make([]int, len(num1)+len(num2))

    // for i:=len(num1)-1; i>=0; i-- {
    //     for j:=len(num2)-1; j>=0; j-- {
    //         // 12 * 34
    //         product := int(num1[i] - '0') * int(num2[j]- '0')

    //         position := i + j + 1
    //         carryPosition := i + j

    //         sum := resultArr[position] + product

    //         resultArr[position] = sum % 10

    //         resultArr[carryPosition]+= sum / 10

    //         //fmt.Println("resultArr:",resultArr) 
    //     }
    // }

    // res := ""
    // idx := 0

    // for idx < len(resultArr) && resultArr[idx] == 0 {
    //     idx++
    // }

    // for i:=idx; i < len(resultArr); i++ {
    //     res += strconv.Itoa(resultArr[i])
    // }


    // return res


    if num1 == "0" || num2 == "0" {
        return "0"
    }


    res := make([]int, len(num1)+len(num2))

    for i:= len(num1)-1; i>=0; i-- {
        for j:= len(num2)-1; j>=0; j-- {
            product := int(num1[i] - '0') * int(num2[j] - '0')

            position := i + j + 1
            carryPosition := i + j

            sum := res[position] + product

            res[position] = sum % 10
            res[carryPosition] += sum / 10 

            //fmt.Println("res:",res)
        }
    }

    idx := 0
    resStr := ""

    //[0 5 6 0 8 8]

    for idx < len(res) && res[idx] == 0 {
        idx++
    }

    for i := idx; i < len(res); i++ {
        resStr += strconv.Itoa(res[i])
    }

    return resStr
}