func largestPalindromic(num string) string {
    var answer string
    numMap := make(map[int]int) 
    for _, v := range num {
        num, _ := strconv.Atoi(string(v))
        numMap[num]++
    }

    for i:=9; i > 0; i-- {
        // if exist
            if numMap[i] % 2 == 1 {
                answer = fmt.Sprintf("%d",i)
                numMap[i]--
                break
            }
    }

    for i:=0; i < 10; i++ {
        if numMap[i] > 0 {
            numMap[i] = numMap[i] / 2 
            str := makeRepeatedStr(i, numMap[i])
            answer = str + answer + str
        }
    }

    answer = strings.Trim(answer, "0")
    if answer == "" {
        return "0"
    }

    return answer
}

func makeRepeatedStr(num, cnt int) string {
    var str string
    for cnt > 0 {
        str += fmt.Sprintf("%d", num)
        cnt--
    }
    return str
}