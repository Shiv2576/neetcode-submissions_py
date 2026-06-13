func maxProduct(nums []int) int {

    n := len(nums)

    if n == 1 {
        return nums[0]
    }


    dp := make([]int , n)

    for i := 0 ; i < n ; i++ {
        max := 0
        prod := 1
        for j := i ; j < n ; j++ {
            prod *= nums[j]

            dp[i] = maximum(prod ,max)

            if max < prod {
                max = prod
            }
        }
    }


    result := maxArr(dp)

    return result


}

func maximum(a, b int) int {
    if a >= b {
        return a
    }

    return b
}


func maxArr(arr []int) int {
    max:= 0

    for _ , num := range arr {
        if num > max {
            max = num
        }
    }

    return max
}
