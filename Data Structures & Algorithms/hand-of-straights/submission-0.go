func isNStraightHand(hand []int, groupSize int) bool {

	if len(hand) % groupSize != 0 {
		return false
	}


	freqmap := make(map[int]int)

	for _ , num := range hand {
		freqmap[num] ++
	}

	sort.Ints(hand)



	for _ , num := range hand {
		if freqmap[num] == 0 {
			continue
		}

		for i:= 0 ; i < groupSize ; i++ {
			next := num + i

			if freqmap[next] == 0 {
				return false
			}

			freqmap[next]--
		}
	}

	return true


    
}
