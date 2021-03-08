package dec10

func countArrangements(adapters []int) int64 {
	compressed := compress(adapters)
	count := int64(1)
	for _, c := range compressed {
		count *= int64(arrangements(c))
	}
	return count
}

// counting the number of elements in each block of ones. 3s are a keeper.
func compress(adapters []int) []int {
	var compressed []int
	var counter int
	for i := range adapters[:len(adapters)-1] {
		dfrce := adapters[i+1] - adapters[i]
		if dfrce == 1 {
			counter++
		} else if dfrce == 3 && counter > 0 {
			compressed = append(compressed, counter)
			counter = 0
		}
	}
	if counter > 0 {
		// don't forget the last one
		compressed = append(compressed, counter)
	}
	return compressed
}

// trib will be used as for memoization on the arrangements method below
var trib = map[int]int{0: 1, 1: 1, 2: 2}

// arrangements is basically a tribonacci
func arrangements(ones int) int {
	a, ok := trib[ones]
	if ok {
		return a
	}
	a = arrangements(ones-1) + arrangements(ones-2) + arrangements(ones-3)
	trib[ones] = a
	return a
}
