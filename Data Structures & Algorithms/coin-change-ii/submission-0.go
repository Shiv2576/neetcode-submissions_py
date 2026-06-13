func change(amount int, coins []int) int {
    m := len(coins)
    n := amount
    dp := make([][]int , m + 1)

    for i:= 0 ; i < m+1 ; i++ {
        dp[i] = make([]int , n+1)
    }

    dp[0][0] = 1

    for i := 1 ; i < n+1 ; i++ {
        dp[0][i] = 0
    }

    for i:= 0 ; i < m+1 ; i++ {
        dp[i][0] = 1
    }


    for i:=1 ; i < m+1 ; i++ {
        coin := coins[i-1]
        for j:=1 ; j < n+1 ; j++ {
            if j >= coin {
                dp[i][j] = dp[i-1][j] + dp[i][j-coin]
            } else {
                dp[i][j] = dp[i-1][j]
            }
        }
    }

    return dp[m][n]
}
