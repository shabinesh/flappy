package main

import (
	"log"
	"path/filepath"

	"github.com/veandco/go-sdl2/sdl"
	sdl_image "github.com/veandco/go-sdl2/sdl_image"
)

type PipePair struct {
	tex         *sdl.Texture
	pipes       []Pipe
	currentPipe int
	speed       int32
}

type Pipe struct {
	cX          int32
	topPipeH    int32
	bottomPipeH int32
}

func (p *PipePair) draw() {
	// TODO draw pipes already generated p.pipes
	// copy to renderer
	renderer.Copy(p.tex, nil, &bottomPipeDst)
	renderer.CopyEx(p.tex, nil, &topPipeDst, 180.0, nil, sdl.FLIP_NONE)
}

func spaceBetweenPipes() (int32, int32) {
	/* generate a space for bird to pass through.
	   randomly generate Y, */
	y := random(30, windowHeight-30)
	return int32(y), int32(y + 100)
}

func (p *PipePair) genPipes(n int) {
	for i := 0; i < n; i++ {
		pipe := Pipe{
			cX: int32(windowWidth),
		}
		pipe.topPipeH, pipe.bottomPipeH = spaceBetweenPipes()
		p.pipes = append(p.pipes, pipe)
	}
}

func NewPipes(n int) *PipePair {
	p := new(PipePair)
	s, err := sdl_image.Load(filepath.Join(mediaPath, "imgs/pipe.png"))
	if err != nil {
		log.Fatal("Failed to load pipe")
	}
	p.tex, err = renderer.CreateTextureFromSurface(s)
	if err != nil {
		log.Fatal("failed to create texture")
	}
	p.speed = 10
	p.genPipes(n)
	return p
}
