package tests

func CompareUnorderedSlice[T comparable](x, y []T) bool {
	if len(x) != len(y) {
		return false
	}

	diff := make(map[T]int, len(x))
	for _, val := range x {
		diff[val]++
	}

	for _, val := range y {
		if _, ok := diff[val]; !ok {
			return false
		}

		diff[val]--
		if diff[val] == 0 {
			delete(diff, val)
		}
	}

	return len(diff) == 0
}
