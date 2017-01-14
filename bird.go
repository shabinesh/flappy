package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Bird struct {
	tex     []*sdl.Texture
	current int
	y       int32
}

func (b *Bird) draw() {
	// TODO Maintain states of each texture being rendered with current
}
