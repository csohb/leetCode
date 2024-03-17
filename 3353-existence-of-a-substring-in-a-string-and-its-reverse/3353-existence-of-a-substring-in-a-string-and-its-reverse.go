func isSubstringPresent(s string) bool {
    // make reverse string 
    // save each substrings with length 2
    // check if contains
    
    var reverse string 
    
    for i:=len(s); i > 0; i-- {
        reverse += string(s[i-1])
    }
    
    
    var substrings []string
    for i:= 0; i < len(s)-1; i++ {
        substr := fmt.Sprintf("%s%s", string(s[i]), string(s[i+1]))
        substrings = append(substrings, substr)
    }
    
    for _, v := range substrings {
        if strings.Contains(reverse, v) {
            return true
        }
    }
    
    return false
}