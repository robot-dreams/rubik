package main

import (
	"fmt"

	"github.com/go-gl/mathgl/mgl32"
)

// The functions in this file "render" an abstract representation of a Rubik's
// Cube into vertex data that can be used by OpenGL for drawing.

func (v ivec3) render() mgl32.Vec3 {
	return mgl32.Vec3{float32(v[0]), float32(v[1]), float32(v[2])}
}

// Given a nonzero input vector u, returns an arbitrary nonzero vector that's
// orthogonal to u.
func orthogonalVec3(u mgl32.Vec3) mgl32.Vec3 {
	for _, axis := range []mgl32.Vec3{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	} {
		v := u.Cross(axis)
		if !v.ApproxEqual(mgl32.Vec3{0, 0, 0}) {
			return v
		}
	}
	panic(fmt.Sprintf("Invalid input vector %v", u))
}

// Given a nonzero input vector u, returns the four corners of a square that's
// orthogonal to u, centered at the origin, and has the given side length.
func orthogonalSquare3(u mgl32.Vec3, l float32) []mgl32.Vec3 {
	v := orthogonalVec3(u).Normalize().Mul(l / 2)
	w := u.Cross(v).Normalize().Mul(l / 2)
	return []mgl32.Vec3{
		v.Add(w),
		v.Sub(w),
		w.Sub(v),
		v.Mul(-1).Sub(w),
	}
}

// A rendered sticker consists of vertex data (position and color) for each of
// the 4 corners of the sticker.  The data layout is as follows:
//
// x1, y1, z1, r1, g1, b1,
// x2, y2, z2, r2, g2, b2,
// ...
//
// In other words, the position offset is 0, the color offset is 12 (3 float32
// values), and the stride is 24 (6 float32 values).
func (s sticker) render(vertexData *[]float32) {
	translation := s.v.render().Add(s.n.render().Mul(0.5))
	// Specifying a side length less than 1 in the call to orthogonalSquare3
	// gives an "exploded cube" look.
	for _, corner := range orthogonalSquare3(s.n.render(), 0.5) {
		vertexPosition := corner.Add(translation)
		*vertexData = append(*vertexData, vertexPosition[:]...)
		*vertexData = append(*vertexData, s.c.rgb()...)
	}
}

// A rendered Rubik's Cube is the concatenation of all rendered stickers,
// together with a slice of element indexes that specify how to group vertices
// into triangles for drawing.
func (r rubiksCube) render() ([]float32, []uint32) {
	vertexData := make([]float32, 0, 24*len(r))
	elementIndexes := make([]uint32, 0, 6*len(r))
	for i, s := range r {
		s.render(&vertexData)
		// The two triangles formed by grouping corners {0, 1, 2} and {1, 2, 3}
		// produce a square.
		//
		// WARNING: These element indexes offsets depend on the order of corners
		// returned by orthogonalSquare3.
		for _, elementIndexOffset := range []uint32{0, 1, 2, 1, 2, 3} {
			elementIndexes = append(
				elementIndexes,
				uint32(4*i)+elementIndexOffset)
		}
	}
	return vertexData, elementIndexes
}
