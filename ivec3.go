package main

// An ivec3 lets us represent vectors in 3D space where the coordinates are
// always integer values.
type ivec3 [3]int

// We pre-define the positive and negative coordinate axes for convenience.
var (
	posX = ivec3{1, 0, 0}
	posY = ivec3{0, 1, 0}
	posZ = ivec3{0, 0, 1}

	negX = ivec3{-1, 0, 0}
	negY = ivec3{0, -1, 0}
	negZ = ivec3{0, 0, -1}
)

func (v ivec3) equal(w ivec3) bool {
	return v[0] == w[0] && v[1] == w[1] && v[2] == w[2]
}

// The following functions rotate the input vector 90 degrees counter-clockwise
// about a given axis.

func rotateX(v ivec3) ivec3 {
	return ivec3{v[0], -v[2], v[1]}
}

func rotateY(v ivec3) ivec3 {
	return ivec3{v[2], v[1], -v[0]}
}

func rotateZ(v ivec3) ivec3 {
	return ivec3{-v[1], v[0], v[2]}
}
