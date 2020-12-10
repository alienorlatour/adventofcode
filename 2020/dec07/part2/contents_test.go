package part2

import (
	"testing"

	"github.com/ablqk/adventofcode/2020/dec07/defs"
	"github.com/stretchr/testify/assert"
)

func TestContainedWiseBRT_Count(t *testing.T) {
	tt := map[string]struct {
		lines   []string
		lookFor defs.Colour
		count   int
	}{
		"two": {
			lines: []string{
				"light red bags contain 1 bright white bag, 2 muted yellow bags.",
			},
			lookFor: "light red",
			count:   3,
		},
		"one": {
			lines: []string{
				"bright white bags contain 1 shiny gold bag.",
			},
			lookFor: "bright white",
			count: 1,
		},
		"no other": {
			lines:   []string{"dotted black bags contain no other bags."},
			lookFor: "dotted black",
			count:   0,
		},
		"a few": {
			lines: []string{
				"light red bags contain 1 bright white bag, 2 muted yellow bags.",
				"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
				"bright white bags contain 1 shiny gold bag.",
				"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
				"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
				"dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
				"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
				"faded blue bags contain no other bags.",
				"dotted black bags contain no other bags.",
			},
			lookFor: "shiny gold",
			count:   32,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			bgt := containedWiseBRT{}
			for _, s := range tc.lines {
				err := bgt.ParseLine(s)
				assert.NoError(t, err)
			}
			assert.Equal(t, tc.count, bgt.Count(tc.lookFor))
		})
	}
}
