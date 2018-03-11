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

func initGlfw(title string) (*glfw.Window, error) {
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
	runtime.LockOSThread()

	window, err := initGlfw("Rubik's Cube")
	if err != nil {
		log.Fatal(err)
	}
	defer glfw.Terminate()

	if err := gl.Init(); err != nil {
		log.Fatal(err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	gl.Enable(gl.DEPTH_TEST)
	// gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)

	program, err := glProgram("vertex.glsl", "fragment.glsl")
	if err != nil {
		log.Fatal(err)
	}
	var vao, vbo, ebo uint32
	gl.GenVertexArrays(1, &vao)
	gl.GenBuffers(1, &vbo)
	gl.GenBuffers(1, &ebo)

	setUniformMatrix4fv(program, "view", mgl32.LookAt(10, 6, 10, 0, 0, 0, 0, 1, 0))
	setUniformMatrix4fv(program, "perspective", mgl32.Perspective(mgl32.DegToRad(30), 1, 0.1, 100))

	r := newRubiksCube()
	r.buffer(vao, vbo, ebo)
	window.SetKeyCallback(r.keyCallback(vao, vbo, ebo))

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		gl.UseProgram(program)
		gl.BindVertexArray(vao)
		gl.DrawElements(gl.TRIANGLES, r.elementCount(), gl.UNSIGNED_INT, nil)
		glfw.PollEvents()
		window.SwapBuffers()
	}
}
