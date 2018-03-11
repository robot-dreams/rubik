package main

import (
	"math"

	"github.com/go-gl/glfw/v3.2/glfw"
)

// 'wasd' controls camera rotation.
func (c *camera) handleRotation(window *glfw.Window, program uint32) {
	if window.GetKey(glfw.KeyA) == glfw.Press {
		c.θ += cameraSpeed
	} else if window.GetKey(glfw.KeyD) == glfw.Press {
		c.θ -= cameraSpeed
	} else if window.GetKey(glfw.KeyW) == glfw.Press {
		// Since ϕ is measured from the y axis, W (which should move the camera
		// upwards) decreases ϕ.
		c.ϕ -= cameraSpeed
		c.ϕ = math.Max(c.ϕ, minϕ)
	} else if window.GetKey(glfw.KeyS) == glfw.Press {
		c.ϕ += cameraSpeed
		c.ϕ = math.Min(c.ϕ, maxϕ)
	} else {
		// Short circuit (and avoid updating the "view" uniform) if the camera
		// wasn't moved.
		return
	}
	setUniformMatrix4fv(program, viewUniform, c.view())
}

// Scroll wheel controls camera zoom.
func (c *camera) zoomCallback(program uint32) glfw.ScrollCallback {
	return func(window *glfw.Window, xOffset, yOffset float64) {
		c.r -= yOffset
		c.r = math.Max(c.r, minR)
		c.r = math.Min(c.r, maxR)
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
