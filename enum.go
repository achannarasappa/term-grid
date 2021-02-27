package grid

type TextAlign int

const (
	Left TextAlign = iota
	Right
)

func (ta TextAlign) String() string {
	return [...]string{"Left", "Right"}[ta]
}

type Overflow int

const (
	Hidden Overflow = iota
	Wrap
	WrapWord
)

func (ta Overflow) String() string {
	return [...]string{"Hidden", "Wrap", "WrapWord"}[ta]
}
