func longestConsecutive(nums []int) int {
    set := make(map[int]bool)

    // put all numbers into set
    for _, num := range nums {
        set[num] = true
    }

    longest := 0

    for num := range set {
        // only start counting if num is the start of a sequence
        if !set[num-1] {
            length := 1
            curr := num

            for set[curr+1] {
                curr++
                length++
            }

            if length > longest {
                longest = length
            }
        }
    }

    return longest
}
