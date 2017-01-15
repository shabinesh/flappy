package main

import (
	"log"
	"path/filepath"

	"github.com/veandco/go-sdl2/sdl"
	sdl_ttf "github.com/veandco/go-sdl2/sdl_ttf"
)

func drawTitle(title string) {
	font, err := sdl_ttf.OpenFont(filepath.Join(mediaPath, "fonts/Flappy.ttf"), 30)
	if err != nil {
		log.Fatal(err, "failed to load font")
	}

	sur, err := font.RenderUTF8_Solid(title, sdl.Color{R: 255, G: 223, B: 0, A: 0})
	if err != nil {
		log.Fatal(err, "failed to create surface")
	}

	tex, err := renderer.CreateTextureFromSurface(sur)
	if err != nil {
		log.Fatal("failed to create texture from surface")
	}
	sur.Free()

	renderer.Copy(tex, nil, &sdl.Rect{X: int32(windowWidth / 2), Y: int32(windowHeight / 2), W: 200, H: 100})
	renderer.Present()
}
