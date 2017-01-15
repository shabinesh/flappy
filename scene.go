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
	s.bird.draw()
	s.pipes.draw()
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

	// Bird
	s.bird = &Bird{}
	for _, f := range []string{"imgs/bird_frame_1.png", "imgs/bird_frame_2.png", "imgs/bird_frame_3.png", "imgs/bird_frame_4.png"} {
		sur, err := sdl_image.Load(filepath.Join(mediaPath, f))
		if err != nil {
			log.Fatal(err, "failed to load bird textures")
		}
		t, err := renderer.CreateTextureFromSurface(sur)
		if err != nil {
			log.Fatal(err, "failed to load texture")
		}
		s.bird.tex = append(s.bird.tex, t)
	}

	s.pipes = NewPipes(3)
	return s
}
