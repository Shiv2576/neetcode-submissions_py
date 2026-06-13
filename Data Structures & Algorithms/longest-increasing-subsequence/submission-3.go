func lengthOfLIS(nums []int) int {

    dp := []int{}


    for _ , num := range nums {
        placed := false


        for i:=0 ; i < len(dp) ; i++ {
            if dp[i] >= num {
                dp[i] = num
                placed = true
                break
            }
        }

        if !placed {
            dp = append(dp , num)
        }
    }

    return len(dp)
}
