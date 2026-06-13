func singleNumber(nums []int) int {
    for i := 0 ; i < len(nums) ; i++ {
        flag := true
        for j:=0 ; j < len(nums) ; j++ {
            if i != j && nums[i] == nums[j] {
                flag = false
                break
            }
        }
        if flag {
            return nums[i]
        }
    }

    return 0

}
