package dec06

func newGroup() group {
	return group{
		yeses: make(map[byte]bool),
	}
}

type group struct {
	yeses    map[byte]bool
	yesCount int
}

func (g *group) InsertAnswer(s []byte) {
	for _, a := range s {
		if !g.yeses[a] {
			g.yeses[a] = true
			g.yesCount++
		}
	}
}

func (g *group) CountYeses() int {
	return g.yesCount
}

// Reset allows pooling and memory reuse
func (g *group) Reset() {
	for k := range g.yeses {
		delete(g.yeses, k)
	}
	g.yesCount = 0
}
