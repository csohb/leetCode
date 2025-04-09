func longestPalindrome(s string) string {
    // sliding window
    // to be palindromic, at least start == end

    left := 0
    right := 0

    maxLen := 0 
    result := ""
for left < len(s) {
        right = left // 각 시작점마다 right을 left로 초기화
        
        // right을 증가시키며 윈도우 크기 확장
        for right < len(s) {
            // 현재 윈도우가 팰린드롬인지 확인
            if checkPalindromic(s[left : right+1]) {
                curLen := right - left + 1
                
                // 더 긴 팰린드롬을 찾았다면 업데이트
                if curLen > maxLen {
                    maxLen = curLen
                    result = s[left : right+1]
                }
            }
            
            right++ // 윈도우 오른쪽 확장
        }
        
        left++ // 다음 시작점으로 이동
    }

    if maxLen == 0 {
        return s[:1]
    }

    return result
}


func checkPalindromic(s string) bool {
    left := 0
    right := len(s) -  1

    for left < right {
        if s[left] != s[right] {
            return false
        }
        left++
        right--
    }
    return true
}