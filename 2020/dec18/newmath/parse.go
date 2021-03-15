package newmath

import (
	"fmt"
	"regexp"
	"strconv"
)

/* See README.md for details on the grammar here. */

// parse an expression
func parse(in string) (expression, error) {
	// only 2 types of starting chars are allowed: numbers and opening parenthesis
	if (in[0] <= '0' || in[0] >= '9') && in[0] != '(' {
		return nil, fmt.Errorf("cannot parse %s, invalid chatracter %c", in, in[0])
	}

	// one scalar
	const regScalar = "^([0-9]+)$"
	r := regexp.MustCompile(regScalar).FindStringSubmatch(in)
	if len(r) == 2 {
		return newScalar(r[1])
	}

	// mul / add
	const regOper = "^(.+) ([*+]) (\\(.+\\)|[0-9]+)$"
	r = regexp.MustCompile(regOper).FindStringSubmatch(in)
	if len(r) != 4 {
		return nil, fmt.Errorf("invalid expression %s, >>>> %d", in, len(r))
	}
	left, err := parse(r[1])
	if err != nil {
		return nil, fmt.Errorf("'invalid left: %w", err)
	}
	right, err := parse(r[3])
	if err != nil {
		return nil, fmt.Errorf("'invalid right: %w", err)
	}

	switch r[2] {
	case "*":
		return multiplication{left, right}, nil
	case "+":
		return addition{left, right}, nil
	}

	// paren
	fmt.Println("TODO : paren !!!!!!") // todo
	return nil, nil
}

// newScalar takes a number and parses it into a valid expression
func newScalar(in string) (expression, error) {
	i, err := strconv.Atoi(in)
	if err != nil {
		return nil, err
	}
	return scalar{i}, err
}
