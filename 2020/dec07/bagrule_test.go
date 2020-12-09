package dec07

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBagRuleTree_ParseLine(t *testing.T) {
	tt := map[string]struct {
		lines        []string
		expectedTree string
	}{
		"two": {
			lines: []string{
				"light red bags contain 1 bright white bag, 2 muted yellow bags.",
			},
			expectedTree: "{map[bright white:[light red] muted yellow:[light red]]}",
		},
		"one": {
			lines: []string{
				"bright white bags contain 1 shiny gold bag.",
			},
			expectedTree: "{map[shiny gold:[bright white]]}",
		},
		"no other": {
			lines:        []string{"dotted black bags contain no other bags."},
			expectedTree: "{map[]}",
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
			expectedTree: "{map[bright white:[light red dark orange] dark olive:[shiny gold] dotted black:[dark olive vibrant plum] faded blue:[muted yellow dark olive vibrant plum] muted yellow:[light red dark orange] shiny gold:[bright white muted yellow] vibrant plum:[shiny gold]]}",
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			bgt := BagRuleTree{}
			for _, s := range tc.lines {
				err := bgt.parseLine(s)
				assert.NoError(t, err)
			}
			assert.Equal(t, tc.expectedTree, fmt.Sprintf("%v", bgt))
		})
	}
}

func TestBagRuleTree_Containers(t *testing.T) {
	bgt := BagRuleTree{}
	for _, s := range []string{
		"light red bags contain 1 bright white bag, 2 muted yellow bags.",
		"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
		"bright white bags contain 1 shiny gold bag.",
		"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
		"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
		"dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
		"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
		"faded blue bags contain no other bags.",
		"dotted black bags contain no other bags."} {
		err := bgt.parseLine(s)
		require.NoError(t, err)
	}

	tt := map[string]struct {
		lookFor      Colour
		expectedSize int
	}{
		"nominal": {
			lookFor:      "shiny gold",
			expectedSize: 4,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			containers := bgt.Containers(tc.lookFor)
			fmt.Println("cc", containers) // todo alienor
			assert.Equal(t, tc.expectedSize, len(containers.Slice()))
		})
	}
}
