package grid_test

import (
	. "github.com/achannarasappa/term-grid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Grid", func() {

	It("should render a grid", func() {
		output := Row(
			50,
			Cell{
				Width:    6,
				Height:   3,
				Text:     "test1 test1 test1 test1",
				Overflow: Wrap,
			},
			Cell{
				Width: 10,
				Text:  "test2",
			},
			Cell{
				Text: "test3",
			},
		)

		// Expect(wrap.String("Hello World!", 7)).To(Equal(""))
		Expect(output).To(Equal(""))
	})

})
