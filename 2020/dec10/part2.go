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

func arrangements(ones int) int {
	// we know from the Internet that groups of ones never go further than 5 elements. There must be a nice mathematical way of solving this, but hey...
	switch ones {
	case 0, 1: return 1
	case 2: return 2
	case 3: return 4
	case 4: return 7
	case 5: return 12
	}
	return 0
}
