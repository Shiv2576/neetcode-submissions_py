func coinChange(coins []int, amount int) int {

    n := amount

    if len(coins) == 1 {
        if n % coins[0] != 0 {
            return -1
        }
    }

    dp := make([]int , n + 1)

    for i := 1 ; i < n+1 ; i++ {
        dp[i] = amount + 1
    }

    dp[0] = 0

    for x := 1 ; x < n + 1 ; x ++ {
        for _ , c := range coins {
            if x >= c {
                dp[x] = min(dp[x] , dp[x-c] + 1)
            }
        }
    }

    if dp[amount] > amount {
        return -1
    }

    return dp[n]
}

func min(a, b int) int {
    if a >= b {
        return b
    } 
    return a
}
