// Package grid provides the functionality to render a configurable text grid
// for use in terminal UI application
package grid

import "strings"

// Grid contains settings that apply to the entire grid
type Grid struct {
	Rows             []Row
	GutterVertical   int
	GutterHorizontal int
}

type gridConfig struct {
	widthGutter int
}

// Renders a grid based on the settings defined by the argument Grid, Row(s), and Cell(s)
// Outputs a string with the settings applied
func Render(grid Grid) string {

	var rows = make([]string, 0)

	config := gridConfig{
		widthGutter: grid.GutterHorizontal,
	}

	for _, row := range grid.Rows {
		rows = append(rows, renderRow(row, config))
	}

	return strings.Join(rows, strings.Repeat("\n", 1+grid.GutterVertical))

}
