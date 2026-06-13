func maxProfit(prices []int) int {
    hold := -prices[0]
    sold := 0
    rest := 0


    for i := 1 ; i < len(prices) ; i++ {
        prevhold := hold
        prevsold := sold
        prevrest := rest

        hold = max(prevhold , prevrest - prices[i])
        sold = prevhold + prices[i]
        rest = max(prevrest , prevsold)
    }

    return max(sold, rest)
    
}


func max(a, b int) int {
    if a >= b {
        return a
    }

    return b
}
