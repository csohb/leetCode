func commonChars(words []string) []string {

    // check map
    // character, how many times
    res := []string{}
    checkMap := make(map[rune]int)

    // check first word 
   

    for _, v := range words[0] {
        if _, ok := checkMap[v]; ok {
            checkMap[v] = checkMap[v] + 1
        } else {
            checkMap[v] = 1
        }
        res = append(res, string(v))
    }

     fmt.Println("checkMap1:",checkMap)
    
    for i:=1; i < len(words); i++ {
        charMap := make(map[rune]int)
        for _, ch := range words[i] {
            if _, ok := charMap[ch]; ok {
                charMap[ch] += 1
            } else {
                charMap[ch] = 1
            }
        }

        fmt.Println("charMap:",charMap)
        for k, _ := range checkMap {
            if checkMap[k] != charMap[k] {
                if checkMap[k] > charMap[k] {
                    checkMap[k] = charMap[k]
                } 
            }
            fmt.Println("checkMap:",checkMap)
        }
    }

    resRes := []string{}

    for k, v := range checkMap {
        if checkMap[k] > 0 {
            for v > 0 {
                resRes = append(resRes, string(k))
                fmt.Println("resRes:",resRes)
                v--
            }
        }
    }

    return resRes
}