package grid

import (
	"math"
	"strings"
)

// Row contains settings for a single row within the parent Grid
type Row struct {
	Width int
	Cells []Cell
}

func getElementWidth(widthTotal int, count int) (int, int) {

	if count <= 0 {
		return widthTotal, 0
	}

	remainder := widthTotal % count
	width := int(math.Floor(float64(widthTotal) / float64(count)))

	return width, remainder
}

func updateCellWidths(widthTotal int, widthGutter int, cells []Cell) []Cell {

	widthFlex := widthTotal
	var widthFlexCells []*int

	for i, cell := range cells {

		if cell.Width <= 0 {
			widthFlexCells = append(widthFlexCells, &cells[i].Width)
		} else {
			widthFlex -= cell.Width
		}

	}

	widthFlex = widthFlex - (widthGutter * (len(cells) - 1))

	widthWithoutRemainder, remainder := getElementWidth(widthFlex, len(widthFlexCells))
	for i := range widthFlexCells {

		*widthFlexCells[i] = widthWithoutRemainder
		if i < remainder {
			*widthFlexCells[i] = widthWithoutRemainder + 1
		}
	}
	return cells

}

func renderRow(row Row, widthGutter int) string {
	lines := []string{}
	heightMax := 0
	widthLinePreviousCells := 0
	cells := updateCellWidths(row.Width, widthGutter, row.Cells)

	for i, cell := range cells {

		if i == (len(cells) - 1) {
			widthGutter = 0
		}

		lines, heightMax = getLines(cell, lines, heightMax, widthLinePreviousCells, widthGutter)

		widthLinePreviousCells += cell.Width + widthGutter
	}

	return strings.Join(lines, "\n")
}
