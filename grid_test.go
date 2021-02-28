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
				GutterVertical:   1,
				GutterHorizontal: 2,
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
								Text:     "test1a test1b test1c test1d",
								Overflow: Wrap,
							},
							{
								Width: 10,
								Text:  "test4a\ntest4b\ntest4c\ntest4d",
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
								Width:    50,
								Text:     "right aligned fixed width cell with word wrap. Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do",
								Align:    Right,
								Overflow: WrapWord,
							},
							{
								Text: "flex width cell here",
							},
						},
					},
					{
						Width: 70,
						Cells: []Cell{
							{
								Width:    35,
								Text:     "fixed width cell that does not take up entire width of row",
								Overflow: WrapWord,
							},
							{
								Width:           20,
								Text:            "cell that is only visible on larger terminals",
								VisibleMinWidth: Large.Size(),
							},
						},
					},
					{
						Width: 70,
						Cells: []Cell{
							{
								Width: 11,
								Text:  "fixed width",
							},
							{
								Text: "flex width 1",
							},
							{
								Text: "flex width 2",
							},
							{
								Text: "flex width 3",
							},
						},
					},
				},
			})

		expected := `

test2 this  test1a tes  test4a      test3 this is another long sentenc
is a long   t1b test1c  test4b                                        
sentence    test1d      test4c                                        
that will               test4d                                        
need to be                                                            
wrapped by                                                            
word                                                                  

    right aligned fixed width cell with word wrap.  flex width cell he
Lorem ipsum dolor sit amet, consectetur adipiscing                    
                                      elit, sed do                    

fixed width cell that does not take
up entire width of row             

fixed width  flex width 1        flex width 2        flex width 3     `

		Expect("\n\n" + output).To(BeIdenticalTo(expected))
	})

	Describe("TextAlign", func() {
		When("left align is selected", func() {
			It("returns the text for left align text", func() {
				input := Left.String()
				output := "Left"
				Expect(input).To(Equal(output))
			})
		})
	})

	Describe("Overflow", func() {
		When("hidden is selected", func() {
			It("returns the text for hidden overflow text", func() {
				input := Hidden.String()
				output := "Hidden"
				Expect(input).To(Equal(output))
			})
		})
	})

	Describe("Breakpoint", func() {
		When("hidden is selected", func() {
			It("returns the breakpoint text", func() {
				input := Small.String()
				output := "Small"
				Expect(input).To(Equal(output))
			})
		})
	})
})
