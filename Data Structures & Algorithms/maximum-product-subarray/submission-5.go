func maxProduct(nums []int) int {

    maxhere := nums[0]
    minhere := nums[0]    
    result := nums[0]


    for i := 1 ; i < len(nums) ; i++ {

        curr := nums[i]

        tempmax := max(curr, max(curr*maxhere , curr*minhere))
        tempmin := min(curr, min(curr*maxhere , curr*minhere))


        maxhere = tempmax
        minhere = tempmin



        result = max(result, maxhere)

    }

    return result
    
}


func max(a, b int) int {
    if a >= b {
        return a
    }

    return b
}

func min(a , b int) int {
    if a >= b {
        return b
    }
    return a
}
