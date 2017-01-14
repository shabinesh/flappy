package main

import (
	sdl_image "github.com/veandco/go-sdl2/sdl_image"

	"log"
	"path/filepath"
)

// Scene has all Scene objects
type Scene struct {
	cam   *Camera
	bird  *Bird
	pipes *PipePair
}

func (s *Scene) drawFrame() {
	s.cam.draw()
}

func NewScene() *Scene {
	s := new(Scene)
	// TODO draw background image
	return s
}
