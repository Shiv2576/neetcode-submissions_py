func checkInclusion(s1 string, s2 string) bool {
    if len(s1) > len(s2) {
        return false
    }

    var need , window [26]int


    for i := 0 ; i < len(s1) ; i++ {

        need[s1[i] - 'a'] ++
        window[s2[i] - 'a'] ++

    }

    if need == window {
        return true
    }

    left := 0

    for right := len(s1) ; right < len(s2) ; right ++ {
        window[s2[right] - 'a'] ++
        window[s2[left] - 'a'] --
        left ++

        if need == window {
            return true
        }
    }

    return false

}
