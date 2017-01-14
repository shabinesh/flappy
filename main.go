package main

import (
	"log"

	"os"

	sdl_ttf "github.com/veandco/go-sdl2/sdl_ttf"

	"github.com/veandco/go-sdl2/sdl"
)

var (
	window        *sdl.Window
	renderer      *sdl.Renderer
	windowHeight  = 600
	windowWidth   = 800
	mediaPath     string
	scene         *Scene
	lastEventTime uint32
)

func init() {
	mediaPath = os.Getenv("FLAPPY_MEDIA")
	if mediaPath == "" {
		log.Fatal("FLAPPY_MEDIA path should be set")
	}
}

func main() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		log.Fatal(err)
	}

	// TODO initialise sdl ttf

	window, renderer, err = sdl.CreateWindowAndRenderer(windowWidth, windowHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		log.Fatal(err)
	}
	window.SetTitle("Golabby")
	drawTitle("GoLab Bird")
	sdl.Delay(3000)
	scene = NewScene()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.WaitEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
				break
			case *sdl.KeyDownEvent:
				handleKeyEvent(event.(*sdl.KeyDownEvent))
			}
		}
	}

	window.Destroy()
	sdl.Quit()
}
