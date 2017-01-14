package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

func handleKeyEvent(e *sdl.KeyDownEvent) {
	lastEventTime = sdl.GetTicks()
}
