package main

import (
	"encoding/binary"
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"math/rand"
	"runtime"
	"time"
)

const (
	gameSizeX, gameSizeY = 128, 128 //Default game size
	cellSizeX, cellSizeY = 5, 5     //Size in pixes of each cell for SDL Window
	chanceOfLive         = 0.2      //On randomize, the chance that each cell is populated
	resetFactor          = 90       //If less than 1/resetFactor cells change in a generation, game will be restart
	fps                  = 25       // Game loop will slow itself down to match target
)

var colorOfCells uint32 = 0x00881111 //Cell Color

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
	fmt.Print("\n")
}

// Draw Game Table to SDL Surface
func (g *game) DrawGame(color uint32, surface *sdl.Surface) {
	// First create a dark background
	bgrect := sdl.Rect{X: 0, Y: 0, W: int32(g.xSize) * cellSizeX, H: int32(g.ySize) * cellSizeY}
	surface.FillRect(&bgrect, 0x11111111)
	for x := 0; x < g.xSize; x++ {
		for y := 0; y < g.ySize; y++ {
			if g.GetCell(x, y) {
				// This cell is alive, draw it
				rect := sdl.Rect{X: int32(x * cellSizeX), Y: int32(y * cellSizeY), W: cellSizeX, H: cellSizeY}
				surface.FillRect(&rect, color)
			}
		}
	}
}

func getRandomColor() uint32 {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 8)
	rand.Read(b)
	return binary.LittleEndian.Uint32(b)
}

// Run Program
func main() {
	var g = NewGame(gameSizeX, gameSizeY)

	rand.Seed(time.Now().UnixNano())
	runtime.LockOSThread()

	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		panic(err)
	}

	defer sdl.Quit()
	window, err := sdl.CreateWindow("Game of Life", 250, 250,
		int32(g.xSize)*cellSizeX, int32(g.ySize)*cellSizeY, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	// Initialize surface we'll be using
	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	running := true

	for running {
		// Begin measuring how long this loop takes
		startTime := time.Now()

		g.DrawGame(colorOfCells, surface)
		window.UpdateSurface()
		g.CreatePlan()
		changed := g.RunPlan()

		// If less than totalCells / resetFactor cells are changed, reinitialize
		if changed < (g.xSize*g.ySize)/resetFactor {
			g.Randomize()
			// Generate random color for every new game
			colorOfCells = getRandomColor()
		}

		// Handle any SDL events that come in
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
				break
			}
		}

		// Check elapsed time and, if necessary, wait for next frame
		if time.Since(startTime) < time.Second/fps {
			time.Sleep((time.Second / fps) - time.Since(startTime))
		}

	}

}
