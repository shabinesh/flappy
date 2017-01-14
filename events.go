package main

import (
	"github.com/veandco/go-sdl2/sdl"

	"fmt"
)

func handleKeyEvent(e *sdl.KeyDownEvent) {
	fmt.Println("Keyboard event occured");
}
