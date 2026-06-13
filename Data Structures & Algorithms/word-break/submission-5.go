func wordBreak(s string, wordDict []string) bool {

    n := len(s)

    dict := make(map[string]bool)

    for _ , w := range wordDict {
        dict[w] = true
    }

    dp := make([]bool, n+1)

    dp[0] = true



    for i := 0 ; i <= n ; i++ {
        for j := 0 ; j < i ; j++ {
            if dp[j] && dict[s[j:i]] {
                dp[i] = true
                break
            }
        }
    }


    return dp[n]
    
}
