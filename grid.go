package main

import (
	"errors"
	"github.com/gdamore/tcell"
)

type grid [][]bool

func newGrid(h, l int) (grid, error) {
	if h < 5 || l < 5 {
		return grid{}, errors.New("error: grid dimensions less than 5")
	}
	g := make([][]bool, h)
	for i := range g {
		g[i] = make([]bool, l)
	}
	return g, nil
}

func (g grid) print(s tcell.Screen, aliveStyle, deadStyle tcell.Style) {
	cellString := "██"
	for y, cells := range g {
		for x, c := range cells {
			if c {
				printmv(s, x*2, y, aliveStyle, cellString)
				continue
			}
			printmv(s, x*2, y, deadStyle, "██")
		}
	}
}

func (g *grid) togggleCell(x, y int) {
	(*g)[y][x] = !(*g)[y][x]
}

func (g *grid) checkNeighbours(x, y int) (int, error) {
	if outOfBounds(*g, x, y) {
		return 0, errors.New("error: coordinates out of bounds")
	}
	directions := map[string][]int{
		"up":          {1, 0},
		"down":        {-1, 0},
		"left":        {0, -1},
		"right":       {0, 1},
		"topLeft":     {-1, -1},
		"topRight":    {-1, 1},
		"bottomLeft":  {1, -1},
		"bottomRight": {1, 1},
	}

	alive := 0

	for _, d := range directions {
		if y+d[0] < 0 || x+d[1] < 0 || y+d[0] > len(*g)-1 || x+d[1] > len((*g)[0])-1 {
			continue
		}

		if (*g)[y+d[0]][x+d[1]] {
			alive++
		}

	}
	return alive, nil
}

func outOfBounds(g grid, x, y int) bool {
	return x < 0 || x > len(g[0])-1 || y < 0 || y > len(g)-1
}

func (g *grid) update() error {
	n, err := g.copyGrid()
	if err != nil {
		return err
	}
	for y, cells := range *g {
		for x, c := range cells {
			a, err := g.checkNeighbours(x, y)
			if err != nil {
				return err
			}

			switch {
			case a < 2 && c:
				n.togggleCell(x, y)
				break

			case a > 3 && c:
				n.togggleCell(x, y)
				break

			case a == 3 && !c:
				n.togggleCell(x, y)
				break
			}
		}
	}
	*g = n
	return nil
}

func (g grid) copyGrid() (grid, error) {
	n, err := newGrid(len(g), len(g[0]))
	if err != nil {
		return nil, err
	}

	for y, cells := range g {
		for x, cell := range cells {
			n[y][x] = cell
		}
	}

	return n, err
}
