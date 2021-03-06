![image](https://github.com/stclaird/golang-sdl2-scrolling/blob/main/githeader.png?raw=true)
# golang-sdl2-scrolling
golang-sdl2-scrolling is repo that uses the SDL2 library and golang to create a scrolling background, much like you might find in a 2D video game.

There are two examples in this repository. The first is a simple scroll effect which moves an image across the screen. The second is slightly more advanced as it moves several image "layers" at different speeds to create what is known as a parallax effect.(https://en.wikipedia.org/wiki/Parallax_scrolling)

## Installation

Choose one of the examples you wish to run and cd to that examples directory within the repo. For example, the instructions will run the parallax example on Linux

```
cd 2_parallax
```

The run the build command

```
go build .
```
This will create a binary called goScrollingParallax

## Running the Demo

Run the command
```
./goScrollingParallax
```
This should if all goes well create a system window and scroll a 3 image layer across the window's view port.
