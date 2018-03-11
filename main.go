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
	vao, vbo, ebo := glObjects()

	s := matstack.NewMatStack()
	s.LeftMul(mgl32.Translate3D(-0.5, -0.5, -0.5))
	s.LeftMul(mgl32.LookAt(10, 6, 10, 0, 0, 0, 0, 1, 0))
	s.LeftMul(mgl32.Perspective(glRadians(30), 1, 0.1, 100))
	mustSetUniformMatrix4fv(program, "transform", s.Peek())

	r := newRubiksCube()
	keyCallback := func(
		window *glfw.Window,
		key glfw.Key,
		_ int,
		action glfw.Action,
		mods glfw.ModifierKey,
	) {
		if action != glfw.Press {
			return
		}
		switch key {
		case glfw.KeyQ:
			r.rotateX(-1)
		case glfw.KeyW:
			r.rotateX(0)
		case glfw.KeyE:
			r.rotateX(1)
		case glfw.KeyA:
			r.rotateY(-1)
		case glfw.KeyS:
			r.rotateY(0)
		case glfw.KeyD:
			r.rotateY(1)
		case glfw.KeyZ:
			r.rotateZ(-1)
		case glfw.KeyX:
			r.rotateZ(0)
		case glfw.KeyC:
			r.rotateZ(1)
		}
	}

	window.SetKeyCallback(keyCallback)
	for !window.ShouldClose() {
		gl.ClearColor(0.1, 0.2, 0.2, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		data, indices := r.render()
		glDraw(program, vao, vbo, ebo, data, indices)
		glDraw(program, vao, vbo, ebo, cubeVertices, cubeIndices)

		glfw.PollEvents()
		window.SwapBuffers()
	}
}
