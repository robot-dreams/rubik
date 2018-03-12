package main

import (
	"fmt"
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

func initGLFW(title string, width, height int) (*glfw.Window, error) {
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
	window, err := initGLFW("Rubik's Cube", 600, 600)
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
	defer gl.DeleteProgram(program)

	var vao, vbo, ebo uint32
	gl.GenVertexArrays(1, &vao)
	gl.GenBuffers(1, &vbo)
	gl.GenBuffers(1, &ebo)
	defer gl.DeleteVertexArrays(1, &vao)
	defer gl.DeleteBuffers(1, &vbo)
	defer gl.DeleteBuffers(1, &ebo)

	perspective := mgl32.Perspective(mgl32.DegToRad(60), 1, 0.1, 100)
	setUniformMatrix4fv(program, perspectiveUniform, perspective)

	c := newCamera()
	setUniformMatrix4fv(program, viewUniform, c.view())
	window.SetScrollCallback(c.zoomCallback(program))

	r := newRubiksCube()
	r.buffer(vao, vbo, ebo)
	window.SetKeyCallback(r.cubeControlCallback(vao, vbo, ebo))

	fmt.Println("Cube control: number keys (1-9)")
	fmt.Println("Camera angle: WASD keys")
	fmt.Println("Camera zoom: mouse scroll")

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
