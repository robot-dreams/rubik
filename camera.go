package main

import (
	"math"

	"github.com/go-gl/mathgl/mgl32"
)

const (
	// We set a range for the distance from the camera to the origin so that the
	// cube still looks reasonable at the min and max allowed values.
	minDistanceToOrigin = 5
	maxDistanceToOrigin = 50
)

type camera struct {
	eye, up mgl32.Vec3
}

func newCamera() *camera {
	return &camera{
		// Just use sensible defaults.
		eye: mgl32.Vec3{0, 0, 10},
		up:  mgl32.Vec3{0, 1, 0},
	}
}

func axisAngleRotation(v mgl32.Vec3, angle float32, axis mgl32.Vec3) mgl32.Vec3 {
	return mgl32.HomogRotate3D(angle, axis).Mul4x1(v.Vec4(1)).Vec3()
}

func (c *camera) adjustPitch(delta float32) {
	axis := c.eye.Cross(c.up).Normalize()
	c.eye = axisAngleRotation(c.eye, delta, axis)
	c.up = axisAngleRotation(c.up, delta, axis).Normalize()
}

func (c *camera) adjustYaw(delta float32) {
	c.eye = axisAngleRotation(c.eye, delta, c.up)
}

func (c *camera) adjustDistanceToOrigin(delta float64) {
	d := float64(c.eye.Len()) + delta

	// Restrict the new distance to the allowed range.
	d = math.Max(d, minDistanceToOrigin)
	d = math.Min(d, maxDistanceToOrigin)

	// Scale the camera position so it has the updated distance.
	c.eye = c.eye.Mul(float32(d) / c.eye.Len())
}

func (c *camera) view() mgl32.Mat4 {
	// The camera always looks towards the origin.
	return mgl32.LookAtV(c.eye, mgl32.Vec3{0, 0, 0}, c.up)
}
