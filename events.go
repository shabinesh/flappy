package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

func handleKeyEvent(e *sdl.KeyDownEvent) {
	lastEventTime = sdl.GetTicks()
	scene.bird.y -= 10
	if scene.bird.gravity < 0 {
		scene.bird.gravity = 0
	}
}
