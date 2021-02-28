package grid

import (
	"strings"

	"github.com/muesli/reflow/ansi"
	"github.com/muesli/reflow/wordwrap"
	"github.com/muesli/reflow/wrap"
)

const (
	fillChar   string = " "
	gutterChar string = " "
)

// Cell contains settings for a single cell within the parent Row
type Cell struct {
	Text     string
	Width    int
	Align    TextAlign
	Overflow Overflow
}

func getText(cell Cell) string {
	if cell.Overflow == Wrap {
		return wrap.String(cell.Text, cell.Width)
	}
	if cell.Overflow == WrapWord {
		return wordwrap.String(cell.Text, cell.Width)
	}
	return cell.Text
}

func getLineText(lineIndex int, lineTexts []string, cellHeightMax int, cellWidth int, cellAlign TextAlign) string {
	if lineIndex < cellHeightMax {
		textWidth := ansi.PrintableRuneWidth(lineTexts[lineIndex])

		if textWidth > cellWidth {
			return lineTexts[lineIndex][:cellWidth]
		}

		if textWidth < cellWidth && cellAlign == Right {
			return strings.Repeat(fillChar, cellWidth-textWidth) + lineTexts[lineIndex]
		}

		if textWidth < cellWidth && cellAlign == Left {
			return lineTexts[lineIndex] + strings.Repeat(fillChar, cellWidth-textWidth)
		}

		return lineTexts[lineIndex]
	}

	return strings.Repeat(fillChar, cellWidth)
}

func getLines(cell Cell, lines []string, heightMax int, widthLinePreviousCells int, widthGutter int) ([]string, int) {

	text := getText(cell)

	textLines := strings.Split(text, "\n")
	cellHeightMax := len(textLines)

	if cellHeightMax > heightMax {
		countLinesToAdd := cellHeightMax - heightMax
		for i := 0; i < countLinesToAdd; i++ {
			lines = append(lines, strings.Repeat(fillChar, widthLinePreviousCells))
		}
		heightMax = cellHeightMax
	}

	for lineIndex := 0; lineIndex < heightMax; lineIndex++ {

		lines[lineIndex] = lines[lineIndex] +
			getLineText(lineIndex, textLines, cellHeightMax, cell.Width, cell.Align) +
			strings.Repeat(gutterChar, widthGutter)

	}

	return lines, heightMax
}
