package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_image"
)

func perror(err error) {
	if err != nil {
		panic(err)
	}
}

// Returns a sdl.Rect that represents the centered src surface on the dst surface
func CalculateCenterRect(src *sdl.Surface, dst *sdl.Surface) (centerRect sdl.Rect) {
	centerRect = src.ClipRect
	centerRect.X = dst.ClipRect.H/2 - centerRect.H/2
	centerRect.Y = dst.ClipRect.W/2 - centerRect.W/2
	return
}

// Blits the src surface with its raw dimensions to the center of the dst surface
// Note : This function is for testing purpose.
func BlitRawCenter(src *sdl.Surface, dst *sdl.Surface) error {
	centerRect := CalculateCenterRect(src, dst)
	return src.Blit(nil, dst, &centerRect)
}

// Blits the src surface with dimensions of srcRect to the center of the dst surface
// Note : This function is for testing purpose
func BlitCenter(srcRect *sdl.Rect, src *sdl.Surface, dst *sdl.Surface) error {
	centerRect := CalculateCenterRect(src, dst)
	return src.Blit(srcRect, dst, &centerRect)
}

// Blits an image to the center of the dst surface
func BlitCenterImage(file string, dst *sdl.Surface) error {
	imgSurface, err := img.Load(file)
	if err != nil {
		return err
	}
	// I don't know what's going on here. Is the resource free at the end of the function ?
	// Is it the intended behaviour ? I should look into it.
	defer imgSurface.Free()
	return BlitRawCenter(imgSurface, dst)
}

func main() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	os.Setenv("LD_LIBRARY_PATH", filepath.Join(dir, "libs"))

	sdl.Init(sdl.INIT_EVERYTHING | img.INIT_PNG)

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)
	perror(err)
	defer window.Destroy()

	// To do : Add handling for events like closing the window.

	// To do : Understand what are renderers
	// renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	// perror(err)

	surface, err := window.GetSurface()
	perror(err)

	BlitCenterImage("res/sdl_img.png", surface)
	window.UpdateSurface()

	sdl.Delay(5000)
	sdl.Quit()
}
