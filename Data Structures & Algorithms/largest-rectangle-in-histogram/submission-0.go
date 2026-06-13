func largestRectangleArea(heights []int) int {

	n := len(heights)
	maxArea := 0

	for i := 0; i < n; i++ {
		length := 0
		breadth := 0
		if heights[i] == 0 {
			continue
		}
		length = heights[i]
		breadth++
		maxArea = max(maxArea, length*breadth)
		for j := i + 1; j < n; j++ {
			if heights[j] == 0 {
				maxArea = max(maxArea, length*breadth)
				break
			} else if heights[j] <= length {
				length = heights[j]
				breadth++
				maxArea = max(maxArea, length*breadth)
			} else {
				breadth++
				maxArea = max(maxArea, length*breadth)
			}
		}
	}

	return maxArea
}
