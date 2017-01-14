Installation Intructions
========================

https://github.com/veandco/go-sdl2



- Init(flags uint32) error
  sdl.SDL_INIT_EVERYTHING init audio, video, timer, events, etc., subsystems
- CreateWindow(title string, x int, y int, w int, h int, flags uint32) (*Window, error)
- CreateWindowAndRenderer(w, h int, flags uint32) (*Window, *Renderer, error)
- sdl.Quit()
- Window.Destroy()

Ref: https://godoc.org/github.com/veandco/go-sdl2/sdl

TODO
 ====

main.go
-------

Create window
Handle events
Quit sdl.
