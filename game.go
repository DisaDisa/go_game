package main

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"
	ttf "github.com/veandco/go-sdl2/ttf"
)

func run() error {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return fmt.Errorf("Couldn't init sdl: %v", err)
	}
	defer sdl.Quit()

	if err := ttf.Init(); err != nil {
		return fmt.Errorf("Couldn't init ttf: %v", err)
	}
	defer ttf.Quit()

	window, render, err := sdl.CreateWindowAndRenderer(800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		return fmt.Errorf("Couldn'r create window")
	}
	defer window.Destroy()

	if err := drawTitle(render); err != nil {
		return fmt.Errorf("Couldn't draw title: %v", err)
	}

	return nil
}

func drawTitle(r *sdl.Renderer) error {
	r.Clear()

	font, err := ttf.OpenFont("/res/IMAGE???", 20)
	if err != nil {
		return fmt.Errorf("Couldn't open font: %v", err)
	}
	defer font.Close()

	col := sdl.Color{R: 255, G: 100, B: 1, A: 1}
	s, err := font.RenderUTF8Solid("Gopher", col)
	if err != nil {
		return fmt.Errorf("Couldn't init title: %v", err)
	}
	defer s.Free()

	texture, err := r.CreateTextureFromSurface(s)
	if err != nil {
		return fmt.Errorf("Couldn't create texture: %v", err)
	}
	defer texture.Destroy()

	if err := r.Copy(texture, nil, nil); err != nil {
		return fmt.Errorf("Copy texture error: %v", err)
	}

	r.Present()
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(2)
	}
}
