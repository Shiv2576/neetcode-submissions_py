func mergeTriplets(triplets [][]int, target []int) bool {
	var found0 , found1 , found2 bool

	for _ , t := range triplets {
		if t[0] > target[0] || t[1] > target[1] || t[2] > target[2] {
			continue
		}

		if t[0] == target[0] {
			found0 = true
		}

		if t[1] == target[1] {
			found1 = true
		}

        if t[2] == target[2] {
			found2 = true
		}
		
	}

	return found0 && found1 && found2
    
}
