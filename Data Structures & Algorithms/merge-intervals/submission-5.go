func merge(intervals [][]int) [][]int {
    n := len(intervals)

    sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

    if n == 1 {
        return intervals
    }
    

    result := [][]int{intervals[0]}

    

    for i := 1 ; i < n ; i++ {

        last := result[len(result) - 1 ]
        curr := intervals[i]

        if last[1] >= curr[0] {
            if curr[1] >= last[1] {
                last[1] = curr[1]
            }

            if last[0] > curr[0] {
                last[0] = curr[0]
            }
        } else {
            result = append(result , intervals[i])
        }
    }

    return result
}
