package main

import (
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

const (
	width  = 600
	height = 600
)

func initGLFW(title string) (*glfw.Window, error) {
	// GLFW must always be called from the same OS thread.
	runtime.LockOSThread()

	if err := glfw.Init(); err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		return nil, err
	}
	window.MakeContextCurrent()

	return window, nil
}

func main() {
	window, err := initGLFW("Rubik's Cube")
	if err != nil {
		log.Fatal(err)
	}
	defer glfw.Terminate()

	if err := initOpenGL(); err != nil {
		log.Fatal(err)
	}

	program, err := newProgram("vertex.glsl", "fragment.glsl")
	if err != nil {
		log.Fatal(err)
	}

	var vao, vbo, ebo uint32
	gl.GenVertexArrays(1, &vao)
	gl.GenBuffers(1, &vbo)
	gl.GenBuffers(1, &ebo)

	perspective := mgl32.Perspective(mgl32.DegToRad(30), 1, 0.1, 100)
	setUniformMatrix4fv(program, perspectiveUniform, perspective)

	c := newCamera(15)
	setUniformMatrix4fv(program, viewUniform, c.view())
	window.SetScrollCallback(c.zoomCallback(program))

	r := newRubiksCube()
	r.buffer(vao, vbo, ebo)
	window.SetKeyCallback(r.cubeControlCallback(vao, vbo, ebo))

	log.Print("Use number keys (1-9) to control the cube")
	log.Print("Use 'wasd' to move the camera")

	for !window.ShouldClose() {
		c.handleRotation(window, program)

		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		gl.UseProgram(program)
		gl.BindVertexArray(vao)
		gl.DrawElements(gl.TRIANGLES, r.elementCount(), gl.UNSIGNED_INT, nil)

		glfw.PollEvents()
		window.SwapBuffers()
	}
}
