package main

import (
	"math"

	"github.com/go-gl/mathgl/mgl32"
)

// We specify the camera position using spherical coordinates.
type camera struct {
	// θ and ϕ are given in radians.
	r, θ, ϕ float64
}

const (
	// We require that ϕ stays between π/12 and 11π/12 so that we can always use
	// (0, 1, 0) as the "up" vector while avoiding degenerate cases.
	minϕ = math.Pi / 12
	maxϕ = 11 * math.Pi / 12

	// We set a range for r so that the cube still looks reasonable at the min
	// and max allowed values.
	minR = 5
	maxR = 50

	// The "camera speed" is an arbitrary value that controls how much θ and ϕ
	// change in response to an input event.
	cameraSpeed = 0.05
)

func newCamera() *camera {
	return &camera{
		// Just use sensible defaults.
		r: 10,
		θ: math.Pi / 4,
		ϕ: math.Pi / 3,
	}
}

func (c *camera) eye() mgl32.Vec3 {
	return mgl32.Vec3{
		float32(c.r * math.Cos(c.θ) * math.Sin(c.ϕ)),
		// Unlike the usual spherical coordinate case, up / down is given by the
		// y coordinate (instead of the z coordinate).
		float32(c.r * math.Cos(c.ϕ)),
		float32(c.r * math.Sin(c.θ) * math.Sin(c.ϕ)),
	}
}

func (c *camera) view() mgl32.Mat4 {
	return mgl32.LookAtV(
		c.eye(),
		// The camera always looks towards the origin.
		mgl32.Vec3{0, 0, 0},
		mgl32.Vec3{0, 1, 0})
}
