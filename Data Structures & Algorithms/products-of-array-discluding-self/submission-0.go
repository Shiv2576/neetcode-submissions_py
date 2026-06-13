func productExceptSelf(nums []int) []int {

    result := []int{}

    for i := range nums {
        result = append(result , product(nums[:i]) * product(nums[i+1:]))
    }

    return result

}


func product(nums []int) int {

    product := 1

    for _ , num := range nums {
        product *= num
    }
    return product
}
