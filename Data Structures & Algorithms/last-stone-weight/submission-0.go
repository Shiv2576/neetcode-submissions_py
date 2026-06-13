func lastStoneWeight(stones []int) int {

    for len(stones) > 1 {
        n := len(stones)

        quickSort(stones, 0 , n-1 )
         
        x := stones[0]
        y := stones[1]

        stones = stones[2:]

        if x != y {
            stones = append(stones, x-y)
        }
    }

    if len(stones) == 0 {
        return 0
    }

    return stones[0]
}


func quickSort(nums []int , low , high int) {
    if low < high {
        p := partitionAsc(nums , low ,high)
        quickSort(nums , low , p-1)
        quickSort(nums , p+1 , high)
    }
}

func partitionAsc(nums []int , low , high int) int {
    pivot := nums[high]
    i := low

    for j := low ; j < high ; j++ {
        if nums[j] >= pivot {
            nums[i] , nums[j] = nums[j] , nums[i]
            i++
        }
    }

    nums[i] , nums[high] = nums[high] , nums[i]

    return i
}