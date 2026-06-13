/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

func canAttendMeetings(intervals []Interval) bool {

    n := len(intervals)

    if n == 0 {
        return true
    }

    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i].end < intervals[j].end
    })

    prevEnd := intervals[0].end

    for i := 1 ; i < n ; i++ {
        currStart := intervals[i].start
        currEnd := intervals[i].end

        if currStart < prevEnd {
            return false
            if currEnd > prevEnd {
                prevEnd = currEnd
            }
        } else {
            prevEnd = currEnd
        }
    }

    return true

}
