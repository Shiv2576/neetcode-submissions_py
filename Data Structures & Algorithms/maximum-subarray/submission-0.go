func maxSubArray(nums []int) int {

    n := len(nums)

    maxSum := nums[0]
    currSum := nums[0]


    for i := 1 ; i < n ; i++ {
        currSum = max(nums[i], currSum + nums[i])
        maxSum = max(maxSum , currSum)
    }

    return maxSum 
}


func max(a, b int) int {
    if a >= b {
        return a
    }
    return b
}