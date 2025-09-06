func reorderSpaces(text string) string {
    var res string

    var spaceCnt int
    strArr := make([]string, 0)
    var word string
    for i, _ := range text {
        if string(text[i]) == " " {
            if word != "" {
                strArr = append(strArr, word)
            }
            word = ""
            spaceCnt++
        } else {
            word += string(text[i])
            if i == len(text) - 1 {
                strArr = append(strArr, word)
            }
        }
    }

    if spaceCnt == 0 {
        return text
    }

    if len(strArr) == 1 {
        res += strArr[0]
        for spaceCnt > 0 {
            res += " "
            spaceCnt--
        }
        return res
    }

    fmt.Println("len(strArr):", len(strArr))

    howMany := spaceCnt / (len(strArr) -1)
    fmt.Println("howMany:", howMany)
    left := spaceCnt % (len(strArr) -1)

    var space string
    for howMany > 0 {
        space += " "
        howMany--
    }

    
    for i, v := range strArr {
        if i == len(strArr) - 1 {
            res += v
            break
        }
        
        res += v + space
        fmt.Println("res:", res)
    }

    if left > 0 {
        for left > 0 {
            res += " "
            left--
        }
    }

    return res
}