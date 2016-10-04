package arithmetic

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Tokenize(input string) ([]Token, error) {

	t := &Tokenizer{
		reader: bufio.NewReader(strings.NewReader(input)),
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
	reader  *bufio.Reader
	payload string
	prev    Token
	output  []Token
	err     error
}

func (t *Tokenizer) push(token Token) {
	t.prev = token
	t.output = append(t.output, token)
}

func (t *Tokenizer) read() (rune, bool) {
	r, _, err := t.reader.ReadRune()
	if err != nil {
		return 0, false
	}

	return r, true
}

func (t *Tokenizer) unread() {
	t.reader.UnreadRune()
}

func startState(t *Tokenizer) stateFunc {
	t.payload = ""

	r, ok := t.read()
	if !ok {
		return nil
	}

	switch {

	case isSpace(r):
		return startState

	case isPlus(r):
		return plusState

	case isMinus(r):
		return minusState

	case isDivide(r):
		return divideState

	case isMultiply(r):
		return multiplyState

	case isModulo(r):
		return moduloState

	case isExponant(r):
		return exponantState

	case isLeftParenthesis(r):
		return leftParenthesisState

	case isRightParenthesis(r):
		return rightParenthesisState

	case isComma(r):
		return commaState

	case isEqual(r):
		return equalState

	case isGreater(r):
		return greaterState

	case isLower(r):
		return lowerState

	case isExclamation(r):
		return exclamationState

	case isAlphaNum(r):
		t.payload += string(r)
		return alphaNumState

	default:
		t.err = fmt.Errorf("unrecognized token: %s", string(r))
		return nil
	}
}

func plusState(t *Tokenizer) stateFunc {
	if t.prev == nil {
		t.push(UnaryPlus{})
		return startState
	}
	operand, operator := t.prev.Value()
	if operand != nil || operator.Kind() == KindRightParenthesis {
		t.push(Plus{})
		return startState
	}
	t.push(UnaryPlus{})
	return startState
}

func minusState(t *Tokenizer) stateFunc {
	if t.prev == nil {
		t.push(UnaryMinus{})
		return startState
	}
	operand, operator := t.prev.Value()
	if operand != nil || operator.Kind() == KindRightParenthesis {
		t.push(Minus{})
		return startState
	}
	t.push(UnaryMinus{})
	return startState
}

func multiplyState(t *Tokenizer) stateFunc {
	t.push(Multiply{})
	return startState
}

func exponantState(t *Tokenizer) stateFunc {
	t.push(Exponant{})
	return startState
}

func divideState(t *Tokenizer) stateFunc {
	t.push(Divide{})
	return startState
}

func moduloState(t *Tokenizer) stateFunc {
	t.push(Modulo{})
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

func equalState(t *Tokenizer) stateFunc {

	r, ok := t.read()
	if !ok {
		t.err = errors.New("unrecognized token \"=\"")
		return nil
	}

	if !isEqual(r) {
		t.err = fmt.Errorf("unrecognized token \"=%s\"", string(r))
		return nil
	}

	t.push(Equal{})
	return startState
}

func greaterState(t *Tokenizer) stateFunc {

	r, ok := t.read()
	if !ok {
		t.err = errors.New("unrecognized token \">\"")
		return nil
	}

	if isEqual(r) {
		t.push(GreaterEqual{})
		return startState
	}
	t.unread()

	t.push(Greater{})
	return startState
}

func lowerState(t *Tokenizer) stateFunc {

	r, ok := t.read()
	if !ok {
		t.err = errors.New("unrecognized token \"<\"")
		return nil
	}

	if isEqual(r) {
		t.push(LowerEqual{})
		return startState
	}
	t.unread()

	t.push(Lower{})
	return startState
}

func exclamationState(t *Tokenizer) stateFunc {

	r, ok := t.read()
	if !ok {
		t.err = errors.New("unrecognized token \"!\"")
		return nil
	}

	if !isEqual(r) {
		t.err = fmt.Errorf("unrecognized token \"!%s\"", string(r))
		return nil
	}

	t.push(Different{})
	return startState
}

func alphaNumState(t *Tokenizer) stateFunc {

	r, ok := t.read()
	if !ok {
		token, err := parse(t.payload)
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
	case isPlus(r):
		fallthrough
	case isMinus(r):
		fallthrough
	case isMultiply(r):
		fallthrough
	case isDivide(r):
		fallthrough
	case isModulo(r):
		fallthrough
	case isExponant(r):
		fallthrough
	case isRightParenthesis(r):
		fallthrough
	case isEqual(r):
		fallthrough
	case isGreater(r):
		fallthrough
	case isLower(r):
		fallthrough
	case isExclamation(r):
		fallthrough
	case isLeftParenthesis(r):
		t.unread()
		token, err := parse(t.payload)
		if err != nil {
			t.err = err
			return nil
		}
		t.push(token)
		return startState

	case isAlphaNum(r):
		t.payload += string(r)
		return alphaNumState

	default:
		t.err = fmt.Errorf("unrecognized token: %s", string(r))
		return nil
	}

}

func isAlphaNum(r rune) bool {
	return unicode.IsDigit(r) || unicode.IsLetter(r) || r == '_' || r == '.'
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

func isEqual(r rune) bool {
	return r == '='
}

func isGreater(r rune) bool {
	return r == '>'
}

func isLower(r rune) bool {
	return r == '<'
}

func isExclamation(r rune) bool {
	return r == '!'
}

func isPlus(r rune) bool {
	return r == '+'
}

func isMinus(r rune) bool {
	return r == '-'
}

func isMultiply(r rune) bool {
	return r == '*'
}

func isDivide(r rune) bool {
	return r == '/'
}

func isModulo(r rune) bool {
	return r == '%'
}

func isExponant(r rune) bool {
	return r == '^'
}

func isSpace(r rune) bool {
	return unicode.IsSpace(r)
}

func parse(input string) (Token, error) {

	input = strings.ToLower(input)

	variable, ok := variables[input]
	if ok {
		return variable, nil
	}

	function, ok := functions[input]
	if ok {
		return function, nil
	}

	for _, exp := range expressions {
		op, ok := exp(input)
		if ok {
			return op, nil
		}
	}

	f, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid numeric value: %s", input)
	}

	return Number(f), nil
}
