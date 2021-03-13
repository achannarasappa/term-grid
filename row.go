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

func getWidthFlexCells(widthTotal int, count int) (int, int) {

	if count <= 0 {
		return widthTotal, 0
	}

	remainder := widthTotal % count
	width := int(math.Floor(float64(widthTotal) / float64(count)))

	return width, remainder
}

func updateCells(widthTotal int, widthGutter int, cells []Cell) []Cell {

	widthFlex := widthTotal
	countFlexWidthCells := 0
	var cellsFiltered = make([]Cell, 0)

	for _, cell := range cells {

		if cell.Width < 0 {
			continue
		}

		if cell.VisibleMinWidth > 0 && cell.VisibleMinWidth > widthTotal {
			continue
		}

		if cell.Width <= 0 {
			countFlexWidthCells++
		}

		widthFlex -= cell.Width
		cellsFiltered = append(cellsFiltered, cell)
	}

	if widthFlex <= 0 {
		return cellsFiltered
	}

	widthFlex = widthFlex - (widthGutter * (len(cellsFiltered) - 1))

	widthFlexCellsWithoutRemainder, widthRemainder := getWidthFlexCells(widthFlex, countFlexWidthCells)

	for i, cell := range cellsFiltered {

		if cell.Width <= 0 && widthRemainder > 0 {
			widthRemainder--
			cellsFiltered[i].Width = widthFlexCellsWithoutRemainder + 1
			continue
		}

		if cell.Width <= 0 {
			cellsFiltered[i].Width = widthFlexCellsWithoutRemainder
			continue
		}

	}

	return cellsFiltered

}

func renderRow(row Row, config gridConfig) string {
	lines := []string{}
	heightMax := 0
	widthLinePreviousCells := 0
	cells := updateCells(row.Width, config.widthGutter, row.Cells)
	cellLastIndex := (len(cells) - 1)

	for i, cell := range cells {

		if i == cellLastIndex {
			config.widthGutter = 0
		}

		lines, heightMax = getLines(cell, lines, heightMax, widthLinePreviousCells, config)

		widthLinePreviousCells += cell.Width + config.widthGutter
	}

	for i, line := range lines {

		if len(line) > row.Width {
			lines[i] = lines[i][:row.Width]
		}

	}

	return strings.Join(lines, "\n")
}
