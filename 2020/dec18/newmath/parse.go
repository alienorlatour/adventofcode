package newmath

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// parse a full expression
func parse(in string) (expression, error) {
	in = removeSpaces(in)
	p := parser{input: in}
	return p.Compile()
}

func removeSpaces(in string) string {
	return strings.ReplaceAll(in, " ", "")
}

/* See README.md for details on the grammar and Turing machine here. */

type parser struct {
	input string
	index int
	tree  expression
	err   error
}

func (p *parser) current() rune {
	if p.index >= len(p.input) {
		return 0
	}
	return rune(p.input[p.index])
}

// Compile the input into an expression
// see the START node of the Turing machine
func (p *parser) Compile() (expression, error) {
	for p.index < len(p.input) {
		if unicode.IsDigit(p.current()) {
			p.parseNumSymb()
		}
		if p.current() == '(' {
			p.parseParens()
		}
	}
	return p.tree, nil
}

func (p *parser) parseNumSymb() {
	start := p.index
	var nbDigits int
	for unicode.IsDigit(p.current()) {
		p.index++
		nbDigits++
	}

	left, err := newScalar(p.input[start:start+nbDigits])
	if err != nil {
		p.err = fmt.Errorf("invalid number: %w", err)
		return
	}

	switch p.current() {
	case '+':
		p.addExpression(addition{left: left})
	case '*':
		p.addExpression(multiplication{left: left})
	case 0:
		// end of the story
		p.addExpression(left)
	default:
		p.err = fmt.Errorf("invalid symbol %q at char %d", p.current(), p.index)
		return
	}
	p.index++
}

func (p *parser) addExpression(sc expression) {
	fmt.Println("adding", sc)
	// todo
	p.tree = sc
}

func (p *parser) parseParens() {
	parCount := 1
	p.index++ // start after the first '('
	start := p.index
	// let's ignore malformed expressions like (()
	for parCount > 0 {
		switch p.current() {
		case '(': parCount++
		case ')': parCount--
		}
		p.index++
	}

	// parse what's inside the parenthesis
	inside := parser{input: p.input[start:p.index-1]}
	exp, err := inside.Compile()
	if err != nil {
		p.err = fmt.Errorf("invalid expression inside parenthesis: %w", err)
	}
	p.addExpression(exp)
}

// newScalar takes a number and parses it into a valid expression
func newScalar(in string) (expression, error) {
	i, err := strconv.Atoi(in)
	if err != nil {
		return nil, err
	}
	return scalar{i}, err
}
