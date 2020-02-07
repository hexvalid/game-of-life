package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"testing"
)

func TestGameTableEquality(t *testing.T) {
	g := NewGame(3, 3)
	g.SetCell(2, 1, true)
	gTarget := NewGame(3, 3)
	gTarget.SetCell(2, 1, true)
	if !g.Equals(gTarget) {
		t.Error("Game tables not equal")
	}
}

// Test 4 Rules of Game
// Reference: https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life#Rules
func TestGameRules(t *testing.T) {
	// RULE 1: Any live cell with fewer than two live neighbours dies, as if by underpopulation.
	g1 := NewGame(3, 3)
	g1.SetCell(1, 1, true)
	g1.SetCell(2, 2, true)
	/*
		0 0 0
		0 1 0
		0 0 1
	*/

	g1Target := NewGame(3, 3)
	/*
		0 0 0
		0 0 0
		0 0 0
	*/

	g1.Run() //not yet implemented
	if !g1.Equals(g1Target) {
		t.Error("RULE 1 failed")
	}

	// RULE 2: Any live cell with two or three live neighbours lives on to the next generation.
	g2 := NewGame(3, 3)
	g2.SetCell(0, 1, true)
	g2.SetCell(1, 1, true)
	g2.SetCell(1, 2, true)
	g2.SetCell(2, 2, true)
	/*
		0 1 0
		0 1 1
		0 0 1
	*/

	g2Target := NewGame(3, 3)
	g2Target.SetCell(0, 1, true)
	g2Target.SetCell(0, 2, true)
	g2Target.SetCell(1, 1, true)
	g2Target.SetCell(1, 2, true)
	g2Target.SetCell(2, 1, true)
	g2Target.SetCell(2, 2, true)
	/*
		0 1 1
		0 1 1
		0 1 1
	*/

	g2.Run() //not yet implemented
	if !g2.Equals(g2Target) {
		t.Error("RULE 2 failed")
	}

	// RULE 3: Any live cell with more than three live neighbours dies, as if by overpopulation.
	g3 := NewGame(3, 3)
	g3.SetCell(0, 0, true)
	g3.SetCell(0, 1, true)
	g3.SetCell(0, 2, true)
	g3.SetCell(1, 0, true)
	g3.SetCell(1, 1, true)
	g3.SetCell(1, 2, true)
	g3.SetCell(2, 0, true)
	g3.SetCell(2, 1, true)
	g3.SetCell(2, 2, true)
	/*
		1 1 1
		1 1 1
		1 1 1
	*/

	g3Target := NewGame(3, 3)
	g3Target.SetCell(0, 0, true)
	g3Target.SetCell(0, 2, true)
	g3Target.SetCell(2, 0, true)
	g3Target.SetCell(2, 2, true)
	/*
		1 0 1
		0 0 0
		1 0 1
	*/

	g3.Run() //not yet implemented
	if !g3.Equals(g3Target) {
		t.Error("RULE 3 failed")
	}

	// RULE 4: Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
	g4 := NewGame(3, 3)
	g4.SetCell(2, 0, true)
	g4.SetCell(2, 1, true)
	g4.SetCell(2, 2, true)
	/*
		0 0 0
		0 0 0
		1 1 1
	*/

	g4Target := NewGame(3, 3)
	g4Target.SetCell(1, 1, true)
	g4Target.SetCell(2, 1, true)
	/*
		0 0 0
		0 1 0
		0 1 0
	*/

	g4.Run() //not yet implemented
	if !g4.Equals(g4Target) {
		t.Error("RULE 4 failed")
	}

}

func TestSDL(t *testing.T) {
	// These are the flags which may be passed to SDL_Init()
	// Reference: https://wiki.libsdl.org/SDL_Init
	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		t.Errorf("SDL is not correctly initialized. Error: %s", err)
	}

	defer sdl.Quit() // Don't forget quit. Prevents memory leak.

	window, err := sdl.CreateWindow("Test", 0, 0, 128, 128, sdl.WINDOW_SHOWN)
	if err != nil {
		t.Errorf("SDL window is not correctly created. Error: %s", err)
	}
	surface, err := window.GetSurface()
	if err != nil {
		t.Errorf("Can't get surface of SDL window. Error: %s", err)
	}

	bg := sdl.Rect{X: 0, Y: 0, W: 128, H: 128}
	err = surface.FillRect(&bg, 0xAAAAAAAA)
	if err != nil {
		t.Errorf("Can't fill surface of SDL window. Error: %s", err)
	}
	err = window.UpdateSurface()
	if err != nil {
		t.Errorf("Can't update surface of SDL window. Error: %s", err)
	}

	rect := sdl.Rect{X: int32(32), Y: int32(32), W: 64, H: 64}
	err = surface.FillRect(&rect, 0x00991111)
	if err != nil {
		t.Errorf("Can't draw rectangle on surface of SDL window. Error: %s", err)
	}
	err = window.UpdateSurface()
	if err != nil {
		t.Errorf("Can't update surface of SDL window. Error: %s", err)
	}

	err = window.Destroy()
	if err != nil {
		t.Errorf("Can't destroy SDL window. Error: %s", err)
	}
}
