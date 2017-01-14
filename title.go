package main

import (
	"log"
	"path/filepath"

	"github.com/veandco/go-sdl2/sdl"
	sdl_ttf "github.com/veandco/go-sdl2/sdl_ttf"
)

func drawTitle(title string) {
	// TODO load title image and display
	renderer.Copy(tex, nil, &sdl.Rect{X: int32(windowWidth / 2), Y: int32(windowHeight / 2), W: 200, H: 100})
	renderer.Present()
}
