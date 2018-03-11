package main

import "github.com/go-gl/glfw/v3.2/glfw"

// The Rubik's Cube can be controlled by pressing the number keys 1-9.  Each of
// the 9 keys rotates a "slice" of the cube 90 degrees counter-clockwise along
// some coordinate axis.
func (r rubiksCube) keyCallback(vao, vbo, ebo uint32) glfw.KeyCallback {
	return func(
		window *glfw.Window,
		key glfw.Key,
		scancode int,
		action glfw.Action,
		mods glfw.ModifierKey,
	) {
		if action != glfw.Press {
			return
		}
		// We rely on the fact that consecutive number keys are specified using
		// consecutive constants.
		switch key {
		case glfw.Key1, glfw.Key2, glfw.Key3:
			r.rotateX(int(key - glfw.Key2))
		case glfw.Key4, glfw.Key5, glfw.Key6:
			r.rotateY(int(key - glfw.Key5))
		case glfw.Key7, glfw.Key8, glfw.Key9:
			r.rotateZ(int(key - glfw.Key8))
		default:
			return
		}
		r.buffer(vao, vbo, ebo)
	}
}
