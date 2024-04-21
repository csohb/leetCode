func reverseWords(s string) string {
    splited:= strings.Split(s, " ")
    var answer string
    fmt.Println("splited length:",len(splited))
    for i:= len(splited) - 1; i >= 0; i-- {
        fmt.Println("splited[i] length:",len(splited[i]))

        if len(splited[i]) == 0 {
            continue
        }
        if i == 0 {
            answer += splited[i]
        } else {
            answer += splited[i] + " "
        }
    }


    last:= strings.TrimSpace(answer)

    return last
}