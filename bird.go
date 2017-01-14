package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Bird struct {
	tex     []*sdl.Texture
	current int
	gravity int32
	y       int32
}

func (b *Bird) draw() {

	// TODO Update gravity and contain bird inside the window

	if b.current >= len(b.tex)-1 {
		b.current = 0
	} else {
		b.current++
	}
	// Size of the bird's texture is 50x50
	renderer.Copy(b.tex[b.current], nil, &sdl.Rect{X: 100, Y: b.y, W: 50, H: 50})
}
