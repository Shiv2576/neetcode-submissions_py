/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

func minMeetingRooms(intervals []Interval) int {
	n := len(intervals)

	starts := make([]int, n)
	ends := make([]int, n)

	for i, it := range intervals {
		starts[i] = it.start
		ends[i] = it.end
	}

	sort.Ints(starts)
	sort.Ints(ends)

	rooms, endPtr := 0, 0

	for i := 0; i < n; i++ {
		if starts[i] < ends[endPtr] {
			rooms++
		} else {
			endPtr++
		}
	}

	return rooms
}

