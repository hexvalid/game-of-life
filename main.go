package main

import (
	"fmt"
	"math/rand"
)

const (
	chanceOfLive = 0.5 //On randomize, the chance that each cell is populated
)

type game struct {
	xSize int    //Horizontal size of game table
	ySize int    //Vertical size of game table
	cells []bool //Contains the active generation of cells
	plans []int  //Stores a plan for the next generation
}

func NewGame(xSize, ySize int) *game {
	return &game{
		xSize: xSize,
		ySize: ySize,
		cells: make([]bool, xSize*ySize),
		plans: make([]int, xSize*ySize),
	}
}

func (g *game) SetCell(xPos, yPos int, val bool) {
	g.cells[yPos*(g.xSize)+xPos] = val
}

func (g *game) GetCell(xPos, yPos int) bool {
	return g.cells[yPos*(g.xSize)+xPos]
}

func (g *game) GetPlan(xPos, yPos int) int {
	return g.plans[yPos*(g.xSize)+xPos]
}

func (g *game) SetPlan(xPos, yPos int, val int) {
	g.plans[yPos*(g.xSize)+xPos] = val
}

// Compare game to another game for unit tests.
func (g *game) Equals(tg *game) bool {
	// Check game tables and cell sizes are equal.
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

// Number of cells adjacent to the target that are alive.
func (g *game) CountNeighbors(x, y int) (neighbors int) {
	up := x - 1
	down := x + 1
	left := y - 1
	right := y + 1

	// cardinal
	if up >= 0 {
		if g.GetCell(up, y) {
			neighbors++
		}
	}
	if down < g.xSize {
		if g.GetCell(down, y) {
			neighbors++
		}
	}
	if left >= 0 {
		if g.GetCell(x, left) {
			neighbors++
		}
	}
	if right < g.ySize {
		if g.GetCell(x, right) {
			neighbors++
		}
	}

	// diagonal
	if up >= 0 && right < g.ySize {
		if g.GetCell(up, right) {
			neighbors++
		}
	}
	if up >= 0 && left >= 0 {
		if g.GetCell(up, left) {
			neighbors++
		}
	}
	if down < g.xSize && right < g.ySize {
		if g.GetCell(down, right) {
			neighbors++
		}
	}
	if down < g.xSize && left >= 0 {
		if g.GetCell(down, left) {
			neighbors++
		}
	}
	return neighbors
}

func (g *game) CreatePlan() {
	for x := 0; x < g.xSize; x++ {
		for y := 0; y < g.ySize; y++ {
			g.SetPlan(x, y, g.CountNeighbors(x, y))
		}
	}
}

func (g *game) RunPlan() (update int) {
	for x := 0; x < g.xSize; x++ {
		for y := 0; y < g.ySize; y++ {
			if g.GetCell(x, y) {
				// This is all about the rules of the game.
				switch g.GetPlan(x, y) {
				case 0, 1:
					// RULE 1: Any live cell with fewer than two live neighbours dies, as if by underpopulation.
					g.SetCell(x, y, false)
					update++
				case 2, 3:
					// RULE 2: Any live cell with two or three live neighbours lives on to the next generation.
				default:
					// RULE 3: Any live cell with more than three live neighbours dies, as if by overpopulation.
					g.SetCell(x, y, false)
					update++
				}
			} else {
				// Determines dead cells.
				switch g.GetPlan(x, y) {
				case 3:
					// RULE 4: Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
					g.SetCell(x, y, true)
					update++
				}
			}
		}
	}
	return update
}

// Randomize populates all cells on game board with a random value.
func (g *game) Randomize() {
	for x := 0; x < g.xSize; x++ {
		for y := 0; y < g.ySize; y++ {
			if rand.Float32() < chanceOfLive {
				g.SetCell(x, y, true)
			} else {
				g.SetCell(x, y, false)
			}
		}
	}
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

// Run Program
func main() {

}
