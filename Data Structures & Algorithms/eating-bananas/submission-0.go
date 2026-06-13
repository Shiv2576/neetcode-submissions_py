func minEatingSpeed(piles []int, h int) int {
    left := 1
    right := maxArr(piles)
    ans := right

    for left <= right {
        mid := (left + right) / 2

        if canEat(piles, h, mid) {
            ans = mid
            right = mid - 1
        } else {
            left = mid + 1
        }
    }

    return ans
}

func canEat(piles []int, h int, k int) bool {
    hours := 0
    for _, p := range piles {
        hours += (p + k - 1) / k
    }
    return hours <= h
} 

func maxArr(arr []int) int {
    max := arr[0]
    for i := 1; i < len(arr); i++ {
        if arr[i] > max {
            max = arr[i]
        }
    }
    return max
}