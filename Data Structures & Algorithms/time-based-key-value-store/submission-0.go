type Pair struct {
	timestamp int
	value string
}


type TimeMap struct {
	store map[string][]Pair

}

func Constructor() TimeMap {
	return TimeMap {
		store : make(map[string][]Pair),
	}

}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.store[key] = append(this.store[key], Pair {
		timestamp : timestamp,
		value : value,
	})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	pairs, ok := this.store[key]
	if !ok {
		return ""
	}

	l , r := 0 , len(pairs) - 1
	res := ""


	for l <= r {
		mid := l + (r - l) / 2

		if pairs[mid].timestamp <= timestamp {
			res = pairs[mid].value
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return res
}
