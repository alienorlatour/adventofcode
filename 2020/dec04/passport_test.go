package dec04

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPassport_IsValid(t *testing.T) {
	tt := map[string]struct {
		p        passport
		expected bool
	}{
		"nominal": {
			p: passport{
				"ecl": "gry",
				"pid": "860033327",
				"eyr": "2020",
				"hcl": "#fffffd",
				"byr": "1937",
				"iyr": "2017",
				"cid": "147",
				"hgt": "183cm",
			},
			expected: true,
		},
		"North Pole": {
			p: passport{
				"ecl": "gry",
				"pid": "860033327",
				"eyr": "2020",
				"hcl": "#fffffd",
				"byr": "1937",
				"iyr": "2017",
				// "cid": missing
				"hgt": "183cm",
			},
			expected: true,
		},
		"invalid": {
			p: passport{
				"ecl": "gry",
				"hcl": "#fffffd",
				"byr": "1937",
				"iyr": "2017",
				"cid": "CID",
				"hgt": "183cm",
			},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.p.IsValid())
		})
	}
}
