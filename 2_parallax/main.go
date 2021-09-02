package main

import (
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 800
	screenHeight = 450
)

const info = `
Application %s starting.
The binary was built with GoLang: %s`

func main() {

	log.Printf(info, "Scrolling Demo", runtime.Version())

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("Init SDL:", err)
		return
	}

	window, err := sdl.CreateWindow(
		"Simple Parallax Scrolling",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight,
		sdl.WINDOW_OPENGL)
	if err != nil {
		fmt.Println("Creating main window:", err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("Creating renderer:", err)
		return
	}
	defer renderer.Destroy()

	bg := initBackground(renderer)
	scroll := true
	
	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}

		renderer.SetDrawColor(22, 26, 26, 255)
		renderer.Clear()

		if scroll == true {
			scroll = scrollBackground(renderer, bg)
		} else {
			time.Sleep(10  * time.Second)
		}
		renderer.Present()
	}

}
