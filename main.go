package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_image"
)

func perror(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	sdl.Init(sdl.INIT_EVERYTHING | img.INIT_PNG)

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)
	perror(err)
	defer window.Destroy()

	// renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	// perror(err)

	surface, err := window.GetSurface()
	perror(err)

	testImgSurface, err := img.Load("res/sdl_img.png")
	perror(err)

	testRect := testImgSurface.ClipRect
	testRect.X = surface.ClipRect.H/2 - testRect.H/2
	testRect.Y = surface.ClipRect.W/2 - testRect.W/2

	testImgSurface.Blit(nil, surface, &testRect)
	window.UpdateSurface()

	sdl.Delay(5000)
	sdl.Quit()
}
