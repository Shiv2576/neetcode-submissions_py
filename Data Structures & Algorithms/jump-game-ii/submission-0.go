func jump(nums []int) int {
	n := len(nums)
	jump := 0
	maxReach := 0
	currentEnd := 0



	for i := 0 ; i < n-1 ; i++ {

		maxReach = max(maxReach , nums[i] + i)

		if i == currentEnd {
			jump ++
			currentEnd = maxReach
		}
	}

	return jump
    
}
