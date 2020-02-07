package main

import "fmt"

type game struct {
	xSize int
	ySize int
	cells []bool
}

func NewGame(xSize, ySize int) *game {
	return &game{
		xSize: xSize,
		ySize: ySize,
		cells: make([]bool, xSize*ySize),
	}
}

func (g *game) SetCell(xPos, yPos int, val bool) {
	g.cells[yPos*(g.xSize)+xPos] = val
}

func (g *game) GetCell(xPos, yPos int) bool {
	return g.cells[yPos*(g.xSize)+xPos]
}

// Compare game to another game for unit tests.
func (g *game) Equals(tg *game) bool {
	// Check game tables and cell sizes are equal
	if g.xSize != tg.xSize || g.ySize != tg.ySize || len(g.cells) != len(tg.cells) {
		return false
	}

	// Check all cell states
	for k := range g.cells {
		if g.cells[k] != tg.cells[k] {
			return false
		}
	}
	return true
}

// Prints game table to terminal. Useful for debugging.
func (g *game) DebugPrint() {
	for x := 0; x < g.xSize; x++ {
		for y := 0; y < g.ySize; y++ {
			if g.cells[y*(g.xSize)+x] {
				fmt.Print("1")
			} else {
				fmt.Print("0")
			}
			fmt.Print(" ")
		}
		fmt.Print("\n")
	}
}

func main() {
	g := NewGame(3, 3)
	g.SetCell(0, 0, true)
	g.DebugPrint()
}
