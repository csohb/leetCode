// func removeStars(s string) string {
//     res := ""
//     for i:=0; i < len(s); i++ {
//         if string(s[i]) == "*" && i - 1 >= 0{
//             res = res[:len(res)-1]
//         } else {
//             res += string(s[i])
//         }
//     }

//     return res
// }

func removeStars(s string) string { 
    // sbyte := []byte(s)

    // stack := []byte{}
    // for i:=0; i<len(sbyte); i++ {
    //     if s[i] == '*' {
    //         stack = stack[:len(stack)-1]
    //     } else {
    //         stack = append(stack, s[i])
    //     }
    // }

    // return string(stack)


    // whenever it meets * then remove the left string

    sByte := []byte{}

    for i:=0; i < len(s); i++ {
        if string(s[i]) == "*" {
            sByte = sByte[:len(sByte)-1]
        } else {
            sByte = append(sByte, s[i])
        }
    }

    return string(sByte)
}