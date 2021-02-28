# term-grid

A simple library to position text and visuals in a grid layout for terminal UIs (TUIs)

## Features

* Responsive - show or hide cells dynamically based on grid size
* Flex layout - dynamically resize cells to use available space
* Align text - align text to the right of left of a cell
* Overflow/Wrap text - set whether text should be hidden, wrap on word, or wrap unconditionally when width exceeds the cell width
* Margins - set width vertically and horizontally between cells

## Install

```sh
go get github.com/achannarasappa/term-grid
```

## Quick Start

```go
// main.go
package main

import (
	"fmt"
	. "github.com/achannarasappa/term-grid"
)

func main() {
	out := Render(
		Grid{
			GutterVertical:   2,
			GutterHorizontal: 5,
			Rows: []Row{
				{
					Width: 100,
					Cells: []Cell{
						{Width: 10, Text: "term-grid is awesome!", Overflow: WrapWord},
						{Width: 10, Text: "everything to the right", Overflow: Wrap, Align: Right},
						{Width: 20, Text: "To Do:\n- take out trash\n- book my flight\n- workout", Overflow: WrapWord},
						{Text: "I'm baby pitchfork iPhone tilde umami man braid"},
					},
				},
			},
		})

	fmt.Print(out)
}
```

```sh
$ go run main.go
term-grid      everything     To Do:                   I'm baby pitchfork iPhone tilde umami man bra
is             to the rig     - take out trash                                                      
awesome!               ht     - book my flight                                                      
                              - workout                                                             
```