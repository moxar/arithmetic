package arithmetic

type operator interface {
	precedence() uint8
	solve(*stack) (interface{}, error)
}
