package arithmetic

type Kind int

const (
	KindFunction Kind = iota
	KindComma
	KindLeftParenthesis
	KindRightParenthesis
	KindOperation
)
