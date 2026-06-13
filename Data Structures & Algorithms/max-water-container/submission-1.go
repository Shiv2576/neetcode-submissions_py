
func maxArea(heights []int) int {

	n := len(heights)

	water := 0
	left := 0
	right := n - 1

	for left < right {
		l := heights[left]
		r := heights[right]

		if l < r {
			water = max(water, l*(right-left))
			left++
		} else if l > r {
			water = max(water, r*(right-left))
			right--
		} else {
			water = max(water, l*(right-left))
			left++
			right--
		}
	}

	return water

}
