package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Grid", func() {
	It("can check the number of neighbours a cell has", func() {
		g, _ := newGrid(5, 5)
		g = grid{
			{false, false, true, false, false},
			{false, true, true, true, false},
			{true, true, false, true, true},
			{false, true, true, true, false},
			{false, false, true, false, false},
		}
		Expect(g.checkNeighbours(1, 0)).To(Equal(3))
	})

	It("will kill off cells with less than 2 neighbours", func() {
		g, _ := newGrid(5, 5)
		g = grid{
			{true, false, false, false, true},
			{false, false, false, false, false},
			{false, false, true, false, false},
			{false, false, false, false, false},
			{true, false, false, false, true},
		}

		_ = g.update()

		Expect(g).To(Equal(grid{
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
		}))
	})

	It("will kill off cells with more than 3 neighbours", func() {
		g, _ := newGrid(5, 5)
		g = grid{
			{true, true, true, true, true},
			{true, true, true, true, true},
			{true, true, true, true, true},
			{true, true, true, true, true},
			{true, true, true, true, true},
		}

		_ = g.update()

		Expect(g).To(Equal(grid{
			{true, false, false, false, true},
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
			{true, false, false, false, true},
		}))
	})

	It("will revive cells with more exactly 3 neighbours", func() {
		g, _ := newGrid(5, 5)
		g = grid{
			{true, true, false, false, false},
			{true, false, false, false, false},
			{false, false, false, false, false},
			{false, false, false, false, true},
			{false, false, false, true, true},
		}

		_ = g.update()

		Expect(g).To(Equal(grid{
			{true, true, false, false, false},
			{true, true, false, false, false},
			{false, false, false, false, false},
			{false, false, false, true, true},
			{false, false, false, true, true},
		}))
	})

	It("will do nothing for patterns other cases", func() {
		g, _ := newGrid(5, 5)
		g = grid{
			{false, false, false, false, false},
			{false, false, true, true, false},
			{false, false, true, true, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
		}

		_ = g.update()

		Expect(g).To(Equal(grid{
			{false, false, false, false, false},
			{false, false, true, true, false},
			{false, false, true, true, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
		}))
	})
})
