package dec02

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDec_Solve(t *testing.T) {
	s, err := New("testdata/in.txt").Solve()
	assert.NoError(t, err)
	assert.Contains(t, s, "15")
}

func Test_points(t *testing.T) {
	tests := []struct {
		name     string
		opponent RPS
		me       RPS
		want     int
	}{
		{name: "RR", opponent: rock, me: rock, want: 3},
		{name: "RP", opponent: rock, me: paper, want: 6},
		{name: "RS", opponent: rock, me: scissors, want: 0},
		{name: "PR", opponent: paper, me: rock, want: 0},
		{name: "PP", opponent: paper, me: paper, want: 3},
		{name: "PS", opponent: paper, me: scissors, want: 6},
		{name: "SR", opponent: scissors, me: rock, want: 6},
		{name: "SP", opponent: scissors, me: paper, want: 0},
		{name: "SS", opponent: scissors, me: scissors, want: 3},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			got := points(tt.opponent, tt.me)
			assert.Equal(t, tt.want, got)
		})
	}
}
