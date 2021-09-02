package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type Layer struct {
	bg_path string  		// path to the image layer
	height  int32   		// height in pixels of this image layer
	width   int32   		// width in pixels of this image layer
	y       float64 		// y cordinate where this image is displayed
	x       float64 		// x cordinate where this image is displayed
	scrollSpeed float64     // speed to scroll this layer (images in the foreground should have a higher scroll speed)
	endScroll int32  		// x co-ordinate at which we stop scrolling this image layer
	texture *sdl.Texture	// image as a n SDL texture

}

func (l *Layer) create(renderer *sdl.Renderer) {
	renderer.Copy(l.texture,
		&sdl.Rect{X: 0, Y: 0, W: l.width, H: l.height},
		&sdl.Rect{X: int32(l.x), Y: int32(l.y), W: l.width, H: l.height})
}

func (l *Layer) scroll() {
	//Move the image 1px 
	l.x -= l.scrollSpeed
}

func getLayers() (layers []Layer) {
	//define background layers. TODO Probably need a way to open image layers from file or filesystem rather than hardcode
	layers = append(layers, Layer{"gfx/0_sky_layer.png", 450, 800, 0, 0, 0, 0, nil})
	layers = append(layers, Layer{"gfx/1_mountain_layer.png", 173, 1450, 200, 0, 0.002, 0, nil})
	layers = append(layers, Layer{"gfx/2_trees_layer.png", 453, 1450, 0, 0, 0.005, 0, nil})

	return layers
}

func initBackground(renderer *sdl.Renderer) (layers []Layer) {
	//initialize background layers

	bg_layers := getLayers() 
	for i, l := range bg_layers {  //cycle through layers
		bgImg, err := img.Load(l.bg_path) //open png from path

		if err != nil {
			fmt.Printf("loading background: %v", err)
		}

		defer bgImg.Free()
		bg_layers[i].texture, err = renderer.CreateTextureFromSurface(bgImg) //create SDL texture from PNG
		if err != nil {
			fmt.Printf("creating background texture: %v", err)
		}

		bg_layers[i].endScroll = (l.width - screenWidth)
	}
	return bg_layers
}

func scrollBackground(renderer *sdl.Renderer, layers []Layer) (scroll bool) {
	//scroll background
	for i, l := range layers {
		l.create(renderer)

		if l.scrollSpeed > 0 {
			if int32(l.x) != (-l.endScroll)  {
				layers[i].scroll()
				scroll = true
			} else {
				scroll = false
			}
		}
	}

	return scroll
}
