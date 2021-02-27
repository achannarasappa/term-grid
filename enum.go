package grid

// TextAlign is an enum used to set the text alignment within a cell
type TextAlign int

const (
	Left TextAlign = iota
	Right
)

func (ta TextAlign) String() string {
	return [...]string{"Left", "Right"}[ta]
}

// Overflow is an enum used to set the text overflow behavior within a cell
type Overflow int

const (
	Hidden Overflow = iota
	Wrap
	WrapWord
)

func (ta Overflow) String() string {
	return [...]string{"Hidden", "Wrap", "WrapWord"}[ta]
}
