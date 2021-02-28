package grid

// TextAlign is an enum used to set the text alignment within a cell
type TextAlign int

const (
	// Left align text within a cell
	Left TextAlign = iota
	// Right align text within a cell
	Right
)

func (ta TextAlign) String() string {
	return [...]string{"Left", "Right"}[ta]
}

// Overflow is an enum used to set the text overflow behavior within a cell
type Overflow int

const (
	// Hide text that extends beyond the width of a cell
	Hidden Overflow = iota
	// Wrap text onto the next line when the text width exceeds the width of the cell
	Wrap
	// Wrap text onto the next line when the text width exceeds the width of the cell
	// but only break on words
	WrapWord
)

func (to Overflow) String() string {
	return [...]string{"Hidden", "Wrap", "WrapWord"}[to]
}

// Breakpoint is an enum used to set the visibility of cells are various terminal sizes
type Breakpoint int

const (
	Small Breakpoint = iota
	Medium
	Large
	XLarge
)

func (bp Breakpoint) String() string {
	return [...]string{"Small", "Medium", "Large", "XLarge"}[bp]
}

func (bp Breakpoint) Size() int {
	return [...]int{50, 100, 200, 400}[bp]
}
