package dec15

type MemoryGame struct {
	// occurences stores the latest
	occurences   map[int]int
	lastSpoken   int
	numberSpoken int
}

func NewMemoryGame(input []int) MemoryGame {
	occurences := make(map[int]int)
	// add to the occurences every value but the last
	for i, n := range input[:len(input)-1] {
		occurences[n] = i + 1
	}
	return MemoryGame{
		occurences:   occurences,
		lastSpoken:   input[len(input)-1],
		numberSpoken: len(input),
	}
}

// ValueAt returns the end-th value spoken
// It is a pointer func, so that the same struct can be reused to compute a higher value without recomputing from scratch
// NOTE: if end is an already-spoken value, you will get the last spoken value - this could be considered a bug or an optimisation of the no-need category
func (m *MemoryGame) ValueAt(end int) int {
	for i := m.numberSpoken; i < end; i++ {
		next := m.occurences[m.lastSpoken]
		if next != 0 {
			next = i - next
		}
		m.occurences[m.lastSpoken] = i
		m.lastSpoken = next
		m.numberSpoken++ // only useful if the func is called again
	}
	return m.lastSpoken
}
