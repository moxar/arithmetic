package arithmetic

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Tokenize(input string) ([]Token, error) {

	t := &Tokenizer{
		input: input,
	}

	for state := startState; state != nil; {
		state = state(t)
	}

	if t.err != nil {
		return nil, t.err
	}

	return t.output, nil
}

type Tokenizer struct {
	input  string
	start  int
	pos    int
	output []Token
	err    error
}

func (t *Tokenizer) push(token Token) {
	t.output = append(t.output, token)
}

func (t *Tokenizer) read() (rune, bool) {
	if len(t.input) < t.start+t.pos {
		return 0, false
	}

	r := t.input[t.start+t.pos]

	t.pos++
	return rune(r), true
}

func (t *Tokenizer) unread() {
	if t.pos > 0 {
		t.pos--
	}
}

func (t *Tokenizer) buffer() string {
	return t.input[t.start : t.start+t.pos]
}

func startState(t *Tokenizer) stateFunc {
	t.start = t.pos
	t.pos = 0

	r, ok := t.read()
	if !ok {
		return nil
	}

	switch {

	case isSpace(r):
		return startState

	case isOperator(r):
		return operatorState

	case isLeftParenthesis(r):
		return leftParenthesisState

	case isRightParenthesis(r):
		return rightParenthesisState

	case isComma(r):
		return commaState

	case isDigit(r):
		fallthrough
	case isLetter(r):
		fallthrough
	case isDot(r):
		return wordState

	default:
		t.err = fmt.Errorf("unrecognized token: %s", string(r))
		return nil
	}
}

func operatorState(t *Tokenizer) stateFunc {

	// NOTE: Unary operator ?
	v := t.buffer()
	var token Token
	switch v {
	case "+":
		token = Plus{}
	case "-":
		token = Minus{}
	case "*":
		token = Multiply{}
	case "/":
		token = Divide{}
	case "%":
		token = Modulo{}
	}

	t.push(token)
	return startState
}

func leftParenthesisState(t *Tokenizer) stateFunc {
	t.push(LeftParenthesis{})
	return startState
}

func rightParenthesisState(t *Tokenizer) stateFunc {
	t.push(RightParenthesis{})
	return startState
}

func commaState(t *Tokenizer) stateFunc {
	t.push(Comma{})
	return startState
}

func wordState(t *Tokenizer) stateFunc {

	r, ok := t.read()
	if !ok {
		token, err := parse(t.buffer())
		if err != nil {
			t.err = err
			return nil
		}
		t.push(token)
		return startState
	}

	switch {

	case isSpace(r):
		fallthrough
	case isComma(r):
		fallthrough
	case isOperator(r):
		fallthrough
	case isRightParenthesis(r):
		fallthrough
	case isLeftParenthesis(r):
		t.unread()
		token, err := parse(t.buffer())
		if err != nil {
			t.err = err
			return nil
		}
		t.push(token)
		return startState

	case isDigit(r):
		fallthrough
	case isLetter(r):
		fallthrough
	case isDot(r):
		return wordState

	default:
		t.err = fmt.Errorf("unrecognized token: %s", string(r))
		return nil
	}

}

func isDigit(r rune) bool {
	return unicode.IsDigit(r)
}

func isLetter(r rune) bool {
	return unicode.IsLetter(r) || r == '_'
}

func isLeftParenthesis(r rune) bool {
	return r == '('
}

func isRightParenthesis(r rune) bool {
	return r == ')'
}

func isComma(r rune) bool {
	return r == ','
}

func isDot(r rune) bool {
	return r == '.'
}

func isOperator(r rune) bool {
	return strings.ContainsRune("/*-+%", r)
}

func isSpace(r rune) bool {
	return unicode.IsSpace(r)
}

func parse(input string) (Token, error) {
	f, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid numeric value: %s", input)
	}

	return Number(f), nil
}
