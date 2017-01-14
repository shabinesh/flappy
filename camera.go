package main

import "github.com/veandco/go-sdl2/sdl"

// Camera ...
type Camera struct {
	tex             *sdl.Texture
	cX              int32
	backgroundWidth int32
	speed           int32
}

func (c *Camera) draw() {
	if c.cX+int32(windowWidth) >= c.backgroundWidth {
		c.cX = 0
	} else {
		c.cX += c.speed
	}
	renderer.Copy(c.tex, &sdl.Rect{X: c.cX, W: int32(windowWidth), H: int32(windowHeight)}, nil)
}
