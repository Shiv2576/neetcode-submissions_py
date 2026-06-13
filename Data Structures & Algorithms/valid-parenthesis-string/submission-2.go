func checkValidString(s string) bool {
	leftStack := []int{}
	starStack := []int{}

	for i, c := range s {
		if c == '(' {
			leftStack = append(leftStack , i)
		} else if c == '*' {
			starStack = append(starStack , i)
		} else {
			if len(leftStack) > 0 {
				leftStack = leftStack[:len(leftStack)-1]
			} else if len(starStack) > 0 {
				starStack = starStack[:len(starStack)-1]
			} else {
				return false
			}
		}
	}

	for len(leftStack) > 0 && len(starStack) > 0 {
		leftIndex := leftStack[len(leftStack) - 1]
		starIndex := starStack[len(starStack) - 1]

		if leftIndex < starIndex {
			leftStack = leftStack[:len(leftStack)-1]
			starStack = starStack[:len(starStack)-1]
		} else {
			break
		}
	}

	return len(leftStack) == 0
    
}
