package main

import (
	"fmt"

	"github.com/go-gl/mathgl/mgl32"
)

func (c color) rgb() []float32 {
	switch c {
	case white:
		return []float32{1, 1, 1}
	case yellow:
		return []float32{0.9, 0.9, 0}
	case orange:
		return []float32{0.9, 0.4, 0.1}
	case red:
		return []float32{0.9, 0, 0}
	case green:
		return []float32{0, 0.7, 0.3}
	case blue:
		return []float32{0, 0, 1}
	default:
		panic(fmt.Sprintf("Unknown color %v", c))
	}
}

func (v ivec3) vec3() mgl32.Vec3 {
	return mgl32.Vec3{float32(v[0]), float32(v[1]), float32(v[2])}
}

var axes3 = []mgl32.Vec3{
	{1, 0, 0},
	{0, 1, 0},
	{0, 0, 1},
}

var zero3 = mgl32.Vec3{0, 0, 0}

func orthogonalVec3(u mgl32.Vec3) mgl32.Vec3 {
	for _, axis := range axes3 {
		v := u.Cross(axis)
		if !v.ApproxEqual(zero3) {
			return v
		}
	}
	return zero3
}

func orthogonalPlane3(u mgl32.Vec3) [4]mgl32.Vec3 {
	v := orthogonalVec3(u).Normalize().Mul(0.48)
	w := u.Cross(v).Normalize().Mul(0.48)
	return [4]mgl32.Vec3{
		v.Add(w),
		v.Sub(w),
		w.Sub(v),
		v.Mul(-1).Sub(w),
	}
}

func (s sticker) render() []float32 {
	corners := orthogonalPlane3(s.n.vec3())
	for i := range corners {
		corners[i] = corners[i].Add(s.v.vec3()).Add(s.n.vec3().Mul(0.5))
	}
	data := make([]float32, 0, 24)
	for _, corner := range corners {
		data = append(data, corner[:]...)
		data = append(data, s.c.rgb()...)
	}
	return data
}

var elementIndices = []uint32{
	0, 1, 2,
	1, 2, 3,
}

func (r rubiksCube) render() ([]float32, []uint32) {
	data := make([]float32, 0, 24*len(r))
	indices := make([]uint32, 0, 6*len(r))
	for i, s := range r {
		data = append(data, s.render()...)
		for _, elementIndex := range elementIndices {
			indices = append(indices, uint32(4*i)+elementIndex)
		}
	}
	return data, indices
}
