func partitionLabels(s string) []int {

	freq := make(map[rune]int)
	result := []int{}


	for i , c := range s {
		freq[c] = i
	}

	size, end := 0 , 0

	for i , c := range s {
		size +=1
		if freq[c] > end {
			end = freq[c]
		}
		if i == end {
			result = append(result , size)
			size = 0
		}
	}

	return result


    
}
