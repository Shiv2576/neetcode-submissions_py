func findTargetSumWays(nums []int, target int) int {
    m := len(nums)

    total := 0
    for _, num := range nums {
        total += num
    }

    if target > total || target < -total {
        return 0
    }

    dp := make([][]int, m+1)
    for i := 0; i <= m; i++ {
        dp[i] = make([]int, 2*total+1)
    }

    dp[0][total] = 1

    for i := 1; i <= m; i++ {
        num := nums[i-1]
        for j := -total; j <= total; j++ {
            prev := dp[i-1][j+total]
            if prev != 0 {
                    dp[i][j+total+num] += prev
                    dp[i][j+total-num] += prev
            }
        }
    }

    return dp[m][target+total]
}
