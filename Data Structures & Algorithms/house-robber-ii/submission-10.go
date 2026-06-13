func rob(nums []int) int {
    n:= len(nums)

    if n == 1 {
        return nums[0]
    }

    nums1 := nums[:n-1]
    nums2 := nums[1:]


    rob1 := maxrob(nums1)
    rob2 := maxrob(nums2)

    return max(rob1, rob2)
}


func maxrob(nums []int) int {
    n:= len(nums)
    dp := make([]int , n)

    switch {
        case n==1:
        return nums[0]
        case n==2:
        return max(nums[0], nums[1])
    }

    dp[0] = nums[0]
    dp[1] = max(nums[0], nums[1])

    for i := 2 ; i < n ; i++ {
        dp[i] = max(dp[i-1] , dp[i-2] + nums[i] , )
    }

    return dp[n-1]
}

func max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}