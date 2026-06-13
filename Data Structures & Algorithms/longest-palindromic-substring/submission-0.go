func longestPalindrome(s string) string {
    
    n := len(s)

    if n == 1 {
        return s
    }

    start , maxlen := 0 , 0 

    for c := 0 ; c < n ; c ++ {

        l , r := c , c

        for l >= 0 && r < n && s[l] == s[r] {
            if r - l +1 > maxlen {
                start = l
                maxlen = r - l + 1
            }
            l --
            r ++
        }

        l , r = c , c+1

        for l >= 0 && r < n && s[l] == s[r] {
            if r - l + 1 > maxlen {
                start = l
                maxlen = r - l + 1
            }
            l -- 
            r ++
        }
    }
    return s[start : start+maxlen]
}
