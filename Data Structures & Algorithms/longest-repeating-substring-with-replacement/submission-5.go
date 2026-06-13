func characterReplacement(s string, k int) int {
    n := len(s)
    freq := make([]int, 26)
    maxFreq := 0
    maxLen := 0
    left := 0


    for right := 0 ; right < n ; right ++ {
        idx := s[right] - 'A'
        freq[idx]++


        if freq[idx] > maxFreq {
            maxFreq = freq[idx]
        }

        windowLen := right - left + 1

        if windowLen - maxFreq > k  {
            leftidx := s[left] - 'A'
            freq[leftidx]--
            left++
        } else {
            maxLen = max(maxLen, windowLen)
        }
    }

    return maxLen

}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
