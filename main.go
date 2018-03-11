package main

import (
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/go-gl/mathgl/mgl32/matstack"
)

const (
	width  = 500
	height = 500
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

	window, err := initGlfw("Cube")
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
	gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)

	program, err := glProgram("vertex.glsl", "fragment.glsl")
	if err != nil {
		log.Fatal(err)
	}
	vao, vbo, ebo := glObjects()

	s := matstack.NewMatStack()
	s.LeftMul(mgl32.Translate3D(-0.5, -0.5, -0.5))
	s.LeftMul(mgl32.LookAt(2, 3, 5, 0, 0, 0, 0, 1, 0))
	s.LeftMul(mgl32.Perspective(glRadians(30), 1, 0.1, 100))
	mustSetUniformMatrix4fv(program, "transform", s.Peek())

	for !window.ShouldClose() {
		gl.ClearColor(0.2, 0.3, 0.3, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		glDraw(program, vao, vbo, ebo, cubeVertices, cubeIndices)

		glfw.PollEvents()
		window.SwapBuffers()
	}
}
