package main

import (
	"log"

	"os"

	"github.com/veandco/go-sdl2/sdl"
)

var (
	window        *sdl.Window
	renderer      *sdl.Renderer
	windowHeight  = 600
	windowWidth   = 800
	mediaPath     string
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

	// TODO Create window

	if err != nil {
		log.Fatal(err)
	}
	window.SetTitle("Golabby")
	sdl.Delay(3000)

	running := true
	for running {
		// TODO Look for events
	}

	window.Destroy()
	sdl.Quit()
}
