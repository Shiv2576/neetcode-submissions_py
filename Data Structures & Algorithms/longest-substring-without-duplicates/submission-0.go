func lengthOfLongestSubstring(s string) int {

    lastseen := make(map[byte]int)
    start := 0 
    maxlength := 0

    for end := 0 ; end < len(s) ; end++ {
        ch := s[end]

        if lastIndex , exist := lastseen[ch]; exist && lastIndex >= start {
            start = lastIndex + 1
        }

        lastseen[ch] = end

        currentlength := end - start + 1
        if currentlength > maxlength {
            maxlength = currentlength
        }
    }

    return maxlength

}
