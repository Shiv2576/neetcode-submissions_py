func groupAnagrams(strs []string) [][]string {
    anagrams := make(map[[26]int][]string)

    for _ , str := range strs {
        count := [26]int{}
        for _ , ch := range str {
            count[ch - 'a']++
        }

        anagrams[count] = append(anagrams[count], str)

    }

    result := make([][]string, 0, len(anagrams))

    for _ , grp := range anagrams {
        result = append(result , grp)

    }

    return result
}
