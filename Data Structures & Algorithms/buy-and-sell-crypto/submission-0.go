func maxProfit(prices []int) int {
    maxprofit := 0 

    for i := 0 ; i < len(prices) ; i++ {
        for j:=i+1 ; j < len(prices) ; j++ {
            diff := prices[j] - prices[i]
            maxprofit = max(maxprofit, diff )
        }
    }

    return maxprofit

}

func max( a, b int) int {

    if a >= b {
        return a
    } else if a <= b {
        return b

    }
    return b
}