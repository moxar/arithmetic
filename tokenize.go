package arithmetic

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// tokenize takes an input and returns an infix slice of tokens.
// note: spaces are not mandatory.
//
// The algorithm is inspired by this:
// https://www.youtube.com/watch?v=HxaD_trXwRE
func tokenize(input string) ([]interface{}, error) {

	t := &tokenizer{
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

// stateFunc describes a function that returns another stateFunc.
// A stateFunc uses its tokenizer to define its output or payload given
// the current or last read rune.
type stateFunc func(t *tokenizer) stateFunc

// tokenizer s a structure that handles a reader and tokenizes it.
type tokenizer struct {
	reader  *bufio.Reader
	payload string
	prev    interface{}
	output  []interface{}
	err     error
}

// push a token to the output.
func (t *tokenizer) push(token interface{}) {
	t.prev = token
	t.output = append(t.output, token)
}

// read the next rune from the reader.
func (t *tokenizer) read() (rune, bool) {
	r, _, err := t.reader.ReadRune()
	if err != nil {
		return 0, false
	}

	return r, true
}

// unread the last read rune of the reader.
func (t *tokenizer) unread() {
	t.reader.UnreadRune()
}

// startState is the default state: it is the state used at the beginning of the input, or after
// a token has been identified.
func startState(t *tokenizer) stateFunc {

	// reset the payload before starting.
	t.payload = ""

	// get the next rune.
	r, ok := t.read()
	if !ok {
		return nil
	}

	// Choose a state given the rune.
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

	case isleftParenthesis(r):
		return leftParenthesisState

	case isrightParenthesis(r):
		return rightParenthesisState

	case isComma(r):
		return commaState

	case isEqual(r):
		return equalState

	case isGreater(r):
		return greaterState

	case isDoubleQuote(r):
		return doubleQuoteState

	case isLower(r):
		return lowerState

	case isAmpersand(r):
		return ampersandState

	case isPipe(r):
		return pipeState

	case isExclamation(r):
		return exclamationState

	case isAlphaNum(r):
		t.payload += string(r)
		return alphaNumState

	default:
		t.err = fmt.Errorf("unrecognized token: \"%s\"", string(r))
		return nil
	}
}

func plusState(t *tokenizer) stateFunc {

	switch t.prev.(type) {
	case rightParenthesis:
		t.push(plus{})
	case float64:
		t.push(plus{})
	case variable:
		t.push(plus{})
	default:
		t.push(unaryPlus{})
	}
	return startState
}

func minusState(t *tokenizer) stateFunc {

	switch t.prev.(type) {
	case rightParenthesis:
		t.push(minus{})
	case float64:
		t.push(minus{})
	default:
		t.push(unaryMinus{})
	}
	return startState
}

func multiplyState(t *tokenizer) stateFunc {
	t.push(multiply{})
	return startState
}

func exponantState(t *tokenizer) stateFunc {
	t.push(exponant{})
	return startState
}

func divideState(t *tokenizer) stateFunc {
	t.push(divide{})
	return startState
}

func moduloState(t *tokenizer) stateFunc {
	t.push(modulo{})
	return startState
}

func leftParenthesisState(t *tokenizer) stateFunc {
	t.push(leftParenthesis{})
	return startState
}

func rightParenthesisState(t *tokenizer) stateFunc {
	t.push(rightParenthesis{})
	return startState
}

func commaState(t *tokenizer) stateFunc {
	t.push(comma{})
	return startState
}

func equalState(t *tokenizer) stateFunc {

	r, ok := t.read()
	if !ok {
		t.err = errors.New("unrecognized token \"=\"")
		return nil
	}

	if !isEqual(r) {
		t.err = fmt.Errorf("unrecognized token \"=%s\"", string(r))
		return nil
	}

	t.push(equal{})
	return startState
}

func greaterState(t *tokenizer) stateFunc {

	r, ok := t.read()
	if !ok {
		t.err = errors.New("unrecognized token \">\"")
		return nil
	}

	if isEqual(r) {
		t.push(greaterEqual{})
		return startState
	}
	t.unread()

	t.push(greater{})
	return startState
}

func lowerState(t *tokenizer) stateFunc {

	r, ok := t.read()
	if !ok {
		t.err = errors.New("unrecognized token \"<\"")
		return nil
	}

	if isEqual(r) {
		t.push(lowerEqual{})
		return startState
	}
	t.unread()

	t.push(lower{})
	return startState
}

func exclamationState(t *tokenizer) stateFunc {

	r, ok := t.read()
	if !ok {
		t.err = errors.New("unrecognized token \"!\"")
		return nil
	}

	if isEqual(r) {
		t.push(different{})
		return startState
	}

	t.unread()
	t.push(not{})
	return startState
}

func doubleQuoteState(t *tokenizer) stateFunc {

	for {
		r, ok := t.read()
		if !ok {
			t.err = fmt.Errorf("expecting \" at end of string: %s", t.payload)
			return nil
		}

		if isDoubleQuote(r) {
			t.push(t.payload)
			return startState
		}

		t.payload += string(r)
	}
}

func ampersandState(t *tokenizer) stateFunc {

	r, ok := t.read()
	if !ok {
		t.err = errors.New("unrecognized token \"&\"")
		return nil
	}

	if isAmpersand(r) {
		t.push(and{})
		return startState
	}
	t.unread()
	t.err = fmt.Errorf("unrecognized tocken \"&%s\"", rune(r))
	return nil
}

func pipeState(t *tokenizer) stateFunc {

	r, ok := t.read()
	if !ok {
		t.err = errors.New("unrecognized token \"|\"")
		return nil
	}

	if isPipe(r) {
		t.push(or{})
		return startState
	}
	t.unread()
	t.err = fmt.Errorf("unrecognized tocken \"|%s\"", rune(r))
	return nil
}

func alphaNumState(t *tokenizer) stateFunc {

	r, ok := t.read()
	if !ok {
		token, err := t.parse()
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
	case isMultiply(r):
		fallthrough
	case isDivide(r):
		fallthrough
	case isModulo(r):
		fallthrough
	case isExponant(r):
		fallthrough
	case isEqual(r):
		fallthrough
	case isGreater(r):
		fallthrough
	case isLower(r):
		fallthrough
	case isExclamation(r):
		fallthrough
	case isDoubleQuote(r):
		fallthrough
	case isAmpersand(r):
		fallthrough
	case isPipe(r):
		fallthrough
	case isrightParenthesis(r):
		fallthrough
	case isleftParenthesis(r):
		fallthrough
	case isMinus(r):
		t.unread()
		token, err := t.parse()
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
		t.err = fmt.Errorf("unrecognized token: \"%s\"", string(r))
		return nil
	}
}

func isAlphaNum(r rune) bool {
	return unicode.IsDigit(r) || unicode.IsLetter(r) || r == '_' || r == '.'
}

func isleftParenthesis(r rune) bool {
	return r == '('
}

func isrightParenthesis(r rune) bool {
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

func isDoubleQuote(r rune) bool {
	return r == '"'
}

func isAmpersand(r rune) bool {
	return r == '&'
}

func isPipe(r rune) bool {
	return r == '|'
}

func isSpace(r rune) bool {
	return unicode.IsSpace(r)
}

// parse the current payload value to tokenize it.
func (t *tokenizer) parse() (interface{}, error) {

	// make input non-case sensitive.
	input := strings.ToLower(t.payload)

	// If the input matches a variable, return it.
	v, ok := variables[input]
	if ok {
		return v, nil
	}

	// If the input matches a function, return it.
	function, ok := functions[input]
	if ok {
		return function, nil
	}

	// If the input matches an alias, return the aliased value.
	a, ok := aliases[input]
	if ok {
		return a, nil
	}

	// If the input can be translated by an expression, return a variable
	// with the input as label, and the expressions output as value.
	for _, exp := range expressions {
		op, ok := exp(input)
		if ok {
			return variable{t.payload, op}, nil
		}
	}

	// If the input can be converted in float, return the numeric value.
	n, err := strconv.ParseFloat(input, 64)
	if err == nil {
		return n, nil
	}

	return nil, fmt.Errorf("unrecognized variable, function or expression: %s", input)
}
