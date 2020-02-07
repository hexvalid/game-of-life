package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"testing"
	"time"
)

func TestSDL(t *testing.T) {
	// These are the flags which may be passed to SDL_Init()
	// Reference: https://wiki.libsdl.org/SDL_Init
	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		t.Errorf("SDL is not correctly initialized. Error: %s", err)
	}
	window, err := sdl.CreateWindow("Test", 0, 0, 128, 128, sdl.WINDOW_SHOWN)
	if err != nil {
		t.Errorf("SDL window is not correctly created. Error: %s", err)
	}
	surface, err := window.GetSurface()
	if err != nil {
		t.Errorf("Can't get surface of SDL window. Error: %s", err)
	}
	time.Sleep(500 * time.Millisecond)
	bg := sdl.Rect{X: 0, Y: 0, W: 128, H: 128}
	err = surface.FillRect(&bg, 0xAAAAAAAA)
	if err != nil {
		t.Errorf("Can't fill surface of SDL window. Error: %s", err)
	}
	err = window.UpdateSurface()
	if err != nil {
		t.Errorf("Can't update surface of SDL window. Error: %s", err)
	}

	time.Sleep(1 * time.Second) // Humanize

	rect := sdl.Rect{X: int32(32), Y: int32(32), W: 64, H: 64}
	err = surface.FillRect(&rect, 0x00991111)
	if err != nil {
		t.Errorf("Can't draw rectangle on surface of SDL window. Error: %s", err)
	}
	err = window.UpdateSurface()
	if err != nil {
		t.Errorf("Can't update surface of SDL window. Error: %s", err)
	}

	time.Sleep(1 * time.Second) //Humanize

	err = window.Destroy()
	if err != nil {
		t.Errorf("Can't destroy SDL window. Error: %s", err)
	}
}
