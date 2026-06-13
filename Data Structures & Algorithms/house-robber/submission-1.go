func rob(nums []int) int {
    n := len(nums)

    switch {
        case n==1:
        return nums[0]
        case n==2:
        return max(nums[0] , nums[1])
        case n==3:
        return max(nums[0]+nums[2] , nums[1])
    }

    dp := make([]int, n)
    dp[0] = nums[0]
    dp[1] = nums[1]
    dp[2] = nums[2] + nums[0]

    for i := 3 ; i < n ; i++ {
        dp[i] = nums[i] + max(dp[i-2], dp[i-3])
    }

    return max(dp[n-1] ,dp[n-2])
    
}

func max(a, b int) int {
    if a >= b {
        return a
    }

    return b
}