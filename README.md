## Installation Intructions
### Install sdl
* fedora: `dnf install SDL2{,_mixer,_image,_ttf}-devel`
* mac: `brew install sdl2{,_image,_ttf,_mixer}`
### Set go environment
* set `$GOPATH` (optional in go1.8)
* `go get -v github.com/shabinesh/flappy`
* `go get -v github.com/veandco/go-sdl2`

### further installations details: [sdl2](https://github.com/veandco/go-sdl2)

## After the exercise:
* you will be able to create Window, renderer
* know how to quit clean.
* handle events

### Functions used: 
* `Init(flags uint32) error` :  example flag `sdl.SDL_INIT_EVERYTHING`: 
* `sdl.CreateWindow(title string, x int, y int, w int, h int, flags uint32) (*Window, error)`
* `sdl.CreateWindowAndRenderer(w, h int, flags uint32) (*Window, *Renderer, error)`
* `sdl.PollEvent()` / `sdl.WaitEvent()`
* `sdl.Quit()`
* `Window.Destroy()`

Ref docs: [https://godoc.org/github.com/veandco/go-sdl2/sdl](https://godoc.org/github.com/veandco/go-sdl2/sdl)

### `main.go`

* Create window
* Handle events
* Quit sdl.
