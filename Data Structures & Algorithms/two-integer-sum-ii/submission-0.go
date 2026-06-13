func twoSum(numbers []int, target int) []int {

    for i := 0 ; i < len(numbers) ; i++ {
        for j := i+1 ; j < len(numbers) ; j++ {
            if numbers[i] + numbers[j] == target && j > i {
                return []int{i+1,j+1}
            }
        }
    }

    return []int{0}
}

