func arrangeWords(text string) string {
    var res string
    text = strings.ToLower(text)

    words := strings.Split(text, " ")

    sort.SliceStable(words, func(i, j int) bool {
        return len(words[i]) < len(words[j])
    })

    res = strings.Join(words, " ")

    res = strings.ToUpper(string(res[0])) + res[1:]

    return res
}