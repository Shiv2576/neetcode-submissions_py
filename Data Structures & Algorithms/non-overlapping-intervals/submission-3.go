func eraseOverlapIntervals(intervals [][]int) int {

	n := len(intervals)

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	fmt.Println(intervals)

	count := 0
	prevEnd := intervals[0][1]

	for i := 1; i < n; i++ {
		currStart := intervals[i][0]
		currEnd := intervals[i][1]
		fmt.Println(prevEnd)
		fmt.Println(currStart, currEnd)

		if currStart < prevEnd {
			count++

			if currEnd < prevEnd {
				prevEnd = currEnd
				fmt.Println(prevEnd)
			}
		} else {
			prevEnd = currEnd
		}

	}

	return count
}
