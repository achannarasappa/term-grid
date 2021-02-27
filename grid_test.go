package grid_test

import (
	. "github.com/achannarasappa/term-grid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Grid", func() {

	It("should render a grid", func() {
		output := Render(
			Grid{
				GutterVertical:   2,
				GutterHorizontal: 5,
				Rows: []Row{
					{
						Width: 70,
						Cells: []Cell{
							{
								Width:    10,
								Text:     "test2 this is a long sentence that will need to be wrapped by word",
								Overflow: WrapWord,
							},
							{
								Width:    10,
								Height:   3,
								Text:     "test1a test1b test1c test1d",
								Overflow: Wrap,
							},
							{
								Width:    10,
								Height:   3,
								Text:     "test4a\ntest4b\ntest4c\ntest4d",
								Overflow: Wrap,
							},
							{
								Text: "test3 this is another long sentence that will need to be wrapped but will overflow",
							},
						},
					},
					{
						Width: 70,
						Cells: []Cell{
							{
								Width:    10,
								Text:     "test2 this is a long sentence that will need to be wrapped by word",
								Overflow: WrapWord,
							},
							{
								Width:    10,
								Height:   3,
								Text:     "test1a test1b test1c test1d",
								Overflow: Wrap,
							},
							{
								Width:    10,
								Height:   3,
								Text:     "test4a\ntest4b\ntest4c\ntest4d",
								Overflow: Wrap,
							},
							{
								Text: "test3 this is another long sentence that will need to be wrapped but will overflow",
							},
						},
					},
				},
			})

		expected := `
test2 this     test1a tes     test4a----     test3 this is another lon
is a long-     t1b test1c     test4b----     -------------------------
sentence--     test1d----     test4c----     -------------------------
that will-     ----------     test4d----     -------------------------
need to be     ----------     ----------     -------------------------
wrapped by     ----------     ----------     -------------------------
word------     ----------     ----------     -------------------------


test2 this     test1a tes     test4a----     test3 this is another lon
is a long-     t1b test1c     test4b----     -------------------------
sentence--     test1d----     test4c----     -------------------------
that will-     ----------     test4d----     -------------------------
need to be     ----------     ----------     -------------------------
wrapped by     ----------     ----------     -------------------------
word------     ----------     ----------     -------------------------`

		Expect("\n" + output).To(BeIdenticalTo(expected))
	})

})
