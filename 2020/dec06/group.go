package dec06

type Group interface {
	InsertAnswer(s []byte)
	CountYeses() int
	Reset()
}

func newGroupPartOne() Group {
	return &groupPrimo{
		yeses: make(map[byte]bool),
	}
}

type groupPrimo struct {
	yeses    map[byte]bool
	yesCount int
}

func (g *groupPrimo) InsertAnswer(s []byte) {
	for _, a := range s {
		if !g.yeses[a] {
			g.yeses[a] = true
			g.yesCount++
		}
	}
}

func (g *groupPrimo) CountYeses() int {
	return g.yesCount
}

// Reset allows pooling and memory reuse
func (g *groupPrimo) Reset() {
	for k := range g.yeses {
		delete(g.yeses, k)
	}
	g.yesCount = 0
}

func newGroupPartTwo() Group {
	return &groupSecundo{
		yeses: make(map[byte]int),
	}
}


// groupSecundo counts the number of yes answers for everyone in the group
// Another algorithm could have been to keep the first answer and then remove all letters that are not present in each of the following answers.
type groupSecundo struct {
	yeses map[byte]int
	totalAnswers int
}

func (g *groupSecundo) InsertAnswer(s []byte) {
	g.totalAnswers++
	for _, a := range s {
		count, ok := g.yeses[a]
		if !ok {
			count = 1
		} else {
			count++
		}
		g.yeses[a] = count
	}
}

func (g *groupSecundo) CountYeses() int {
	var count int
	for _, c := range g.yeses {
		if c == g.totalAnswers {
			count++
		}
	}
	return count
}

func (g *groupSecundo) Reset() {
	for k := range g.yeses {
		delete(g.yeses, k)
	}
	g.totalAnswers = 0
}
