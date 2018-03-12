package main

import (
	"github.com/go-gl/glfw/v3.2/glfw"
)

const (
	// The "camera speed" is an arbitrary value that controls how much the
	// camera moves in response to an input event.
	cameraSpeed = 0.05
)

// WASD keys control camera rotation.
func (c *camera) handleRotation(window *glfw.Window, program uint32) {
	if window.GetKey(glfw.KeyW) == glfw.Press {
		c.adjustPitch(cameraSpeed)
	} else if window.GetKey(glfw.KeyA) == glfw.Press {
		c.adjustYaw(cameraSpeed)
	} else if window.GetKey(glfw.KeyS) == glfw.Press {
		c.adjustPitch(-cameraSpeed)
	} else if window.GetKey(glfw.KeyD) == glfw.Press {
		c.adjustYaw(-cameraSpeed)
	} else {
		// Short circuit (and avoid updating the "view" uniform) if the camera
		// wasn't moved.
		return
	}
	setUniformMatrix4fv(program, viewUniform, c.view())
}

// Mouse scroll controls camera zoom.
func (c *camera) zoomCallback(program uint32) glfw.ScrollCallback {
	return func(window *glfw.Window, xOffset, yOffset float64) {
		c.adjustDistanceToOrigin(-yOffset)
		setUniformMatrix4fv(program, viewUniform, c.view())
	}
}

// Number keys (1-9) control the Rubik's Cube.  Each key rotates some "slice" of
// the cube 90 degrees counter-clockwise along some coordinate axis.
func (r rubiksCube) cubeControlCallback(vao, vbo, ebo uint32) glfw.KeyCallback {
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
