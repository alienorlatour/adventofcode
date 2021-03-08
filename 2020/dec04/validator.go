package dec04

import (
	"regexp"
	"strconv"
)

// EntryValidator returns whether a passport entry is valid
type EntryValidator func(string) bool

// mandatoryKeys list the mandatory keys and how to validate each value
var mandatoryKeys = map[string]EntryValidator{
	"byr": yearValidator(1920, 2002),
	"iyr": yearValidator(2010, 2020),
	"eyr": yearValidator(2020, 2030),
	"hgt": hgtValidator,
	"hcl": regexpValidator("#[0-9a-f]{6}"),
	"ecl": regexpValidator("amb|blu|brn|gry|grn|hzl|oth"),
	"pid": regexpValidator("[0-9]{9}"),
}

// yearValidator returns an EntryValidator for years between 2 values (included)
func yearValidator(min, max int) func(string) bool {
	return func(s string) bool {
		y, err := strconv.Atoi(s)
		if err != nil {
			return false
		}
		return y >= min && y <= max
	}
}

// regexpValidator returns an EntryValidator for a given string. The string should match the parameter completely (no trailing characters)
func regexpValidator(reg string) func(string) bool {
	return func(s string) bool {
		r := regexp.MustCompile("^" + reg + "$")
		return r.Match([]byte(s))
	}
}

// hgtValidator validates the height, because it's too complicated. Children and small people have no right to fly.
func hgtValidator(s string) bool {
	r := regexp.MustCompile("([0-9]+)(in|cm)")
	if !r.Match([]byte(s)) {
		return false
	}
	expression := r.FindStringSubmatch(s)
	h, err := strconv.Atoi(expression[1])
	if err != nil {
		return false
	}
	switch expression[2] {
	case "cm":
		return h >= 150 && h <= 193
	case "in":
		return h >= 59 && h <= 76
	default:
		return false
	}
}
