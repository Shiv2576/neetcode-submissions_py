func countSubstrings(s string) int {
    	n := len(s)
	if n == 1 {
		return 1
	}

	maxLen := 1 
	count := 0

	for c := 0; c < n; c++ {
		// -------- Odd-length palindromes --------
		l, r := c, c
		for l >= 0 && r < n && s[l] == s[r] {
			count++
			if r-l+1 > maxLen {
				maxLen = r - l + 1
			}
			l--
			r++
		}

		// -------- Even-length palindromes --------
		l, r = c, c+1
		for l >= 0 && r < n && s[l] == s[r] {
			count++
			if r-l+1 > maxLen {
				maxLen = r - l + 1
			}
			l--
			r++
		}
	}

	return count
}
