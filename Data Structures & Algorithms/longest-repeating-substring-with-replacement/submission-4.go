func characterReplacement(s string, k int) int {
    n := len(s)
    freq := make([]int, 26) // Frequency array for uppercase letters
    maxFreq := 0
    maxLen := 0
    left := 0
    
    for right := 0; right < n; right++ {
        // Add current character to window
        idx := s[right] - 'A'
        freq[idx]++
        
        // Update max frequency in current window
        if freq[idx] > maxFreq {
            maxFreq = freq[idx]
        }
        
        // Check if window is valid
        windowLen := right - left + 1
        if windowLen - maxFreq > k {
            // Invalid window, shrink from left
            leftIdx := s[left] - 'A'
            freq[leftIdx]--
            left++
        } else {
            // Valid window, update max length
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