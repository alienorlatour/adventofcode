package dec03

type forest struct {
	carto [][]byte
}

func (f forest) countSlope(sl slope) int {
	var count int
	// 0,0 is the starting position but is should not be counted
	i := sl.down
	j := sl.right
	w := len(f.carto[0]) // the forest is a rectangle

	for i < len(f.carto) {
		if f.carto[i][j] == treeByte {
			count++
		}
		i += sl.down
		j = (j + sl.right) % w // this is Glastonbury all over again!
	}
	return count
}
