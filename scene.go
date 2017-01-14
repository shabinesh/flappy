package main

import (
	sdl_image "github.com/veandco/go-sdl2/sdl_image"

	"log"
	"path/filepath"
)

// Scene has all Scene objects
type Scene struct {
	cam *Camera
}

func (s *Scene) drawFrame() {
	s.cam.draw()
}

func NewScene() *Scene {
	s := new(Scene)
	sur, err := sdl_image.Load(filepath.Join(mediaPath, "imgs/background.png"))
	if err != nil {
		log.Fatal(err, "failed to load background")
	}
	s.cam = &Camera{}

	s.cam.tex, err = renderer.CreateTextureFromSurface(sur)
	if err != nil {
		log.Fatal(err, "failed to load texture")
	}
	sur.Free()
	s.cam.speed = 10
	s.cam.backgroundWidth = 2000

	return s
}
