package grid

import "strings"

type Grid struct {
	Rows             []Row
	GutterVertical   int
	GutterHorizontal int
}

func Render(grid Grid) string {

	var rows = make([]string, 0)

	for _, row := range grid.Rows {
		rows = append(rows, renderRow(row, grid.GutterHorizontal))
	}

	return strings.Join(rows, strings.Repeat("\n", 1+grid.GutterVertical))

}
