package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type background struct {
	texure    *sdl.Texture
	x         float64
	y         float64
	height    int32
	width     int32
	offScreen int32
}

const scrollSpeed float64 = 0.005

func newBackground(renderer *sdl.Renderer) (b background, err error) {

	bgImg, err := img.Load("gfx/background1.png")
	if err != nil {
		return background{}, fmt.Errorf("loading background: %v", err)
	}
	defer bgImg.Free()
	b.texure, err = renderer.CreateTextureFromSurface(bgImg)
	if err != nil {
		return background{}, fmt.Errorf("creating background texture: %v", err)
	}

	//retrieve image attributes
	_, _, imageWidth, imageHeight, err := b.texure.Query()
	b.height = imageHeight
	b.width = imageWidth
	b.offScreen = (b.width - 800)

	return b, nil
}

func (b *background) create(renderer *sdl.Renderer) {
	renderer.Copy(b.texure,
		&sdl.Rect{X: 0, Y: 0, W: b.width, H: b.height},
		&sdl.Rect{X: int32(b.x), Y: int32(b.y), W: b.width, H: b.height})
}

func (b *background) scroll() {

	if int32(b.x) != (-b.offScreen) {
		b.x -= scrollSpeed
	}

}
