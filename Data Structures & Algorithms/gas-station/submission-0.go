func canCompleteCircuit(gas []int, cost []int) int {

	total := 0 
	tank := 0
	start := 0

	for i := start ; i < len(gas) ; i++ {
		gain := gas[i] - cost[i]
		total += gain
		tank += gain

		if tank < 0 {
			start = i + 1
			tank = 0
		}
	}


	if total < 0  {
		return -1
	} else {
		return start
	}
}
