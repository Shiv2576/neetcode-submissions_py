func canPartition(nums []int) bool {

	sort.Slice(nums, func(a, b int) bool {
		return nums[a] < nums[b]
	})

	fmt.Println(nums)

	totalSum := sum(nums)

	if totalSum%2 != 0 {
		return false
	}

	target := totalSum / 2
	dp := make([]bool, target+1)
	dp[0] = true

	for _, num := range nums {
		for s := target; s >= num; s-- {
			dp[s] = dp[s] || dp[s-num]
		}
	}

	return dp[target]
}

func sum(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}

	return sum
}