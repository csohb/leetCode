func wordBreak(s string, wordDict []string) bool {
    dp := make([]bool, len(s)+1)
    dp[0] = true

    for i, reachable := range dp {
        if reachable == false {
            continue
        }

        for _, word := range wordDict {
            to := i + len(word)
            if to > len(s) {
                continue
            }

            if s[i:to] == word {
                dp[to] = true
            }
        }
    }
    return dp[len(s)]
}