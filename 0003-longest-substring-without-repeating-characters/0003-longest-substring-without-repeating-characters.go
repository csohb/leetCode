func lengthOfLongestSubstring(s string) int {
    // sliding window
    max := 0
    charMap := make(map[byte]int)
    left := 0
    right := 0
    for right < len(s) {
        charMap[s[right]]++

        for left < right && charMap[s[right]] > 1 {
            charMap[s[left]]--
            left++
        }

        if max < (right - left + 1) {
            max = right - left + 1
        }
        right++
    }

    return max
}
