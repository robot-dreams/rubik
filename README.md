# rubik

This repository contains a Rubik's Cube implementation using go and OpenGL.  Both the cube and the camera can be controlled by user input.  Note that this is purely a graphics program; there's no features like automatic solving.

[Here's how the cube looks.](https://github.com/robot-dreams/rubik/blob/master/screenshot.png)

## Controls

- Cube controls: number keys (1-9)
- Camera angle: WASD keys
- Camera zoom: mouse scroll

## Instructions

If you don't have go installed yet, you can [download the latest version of go from the official website](https://golang.org/dl/).  On Mac OS X, you will also need a recent version of Xcode and the command line tools.

```
# Fetch dependencies.
go get github.com/go-gl/gl/v4.1-core/gl
go get github.com/go-gl/glfw/v3.2/glfw
go get github.com/go-gl/mathgl/mgl32

# Clone repo and build everything.
git clone https://github.com/robot-dreams/rubik.git; cd rubik; go build

# Start the program!
./rubik
```

## Implementation Notes

- The code for handling the Rubik's Cube itself is in `sticker.go` and `rubik.go`
- Most of the "interesting" graphics code is in `render.go` and `camera.go`
- User input is, unsurprisingly, handled in `input.go`

## Potential Extensions

- [x] Unrestricted camera rotation
- [x] Basic lighting
- [ ] Animated cube operations
- [ ] Configurable cube size (e.g. 2x2, 5x5)
- [ ] Cube controls using the mouse
- [ ] Improved appearance (e.g. multiple light sources, fog, neon, opacity)
