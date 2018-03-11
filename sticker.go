package main

// A Rubik's Cube can be represented in terms of individual stickers.
type sticker struct {
	c color

	// Each sticker is stuck to a mini-cube; v is the position of the center of
	// the mini-cube.
	//
	// Every component of v must be in {-1, 0, 1}.
	v ivec3

	// Each sticker faces some direction; n is a unit vector in that direction.
	//
	// n must be one of {posX, posY, posZ, negX, negY, negZ}.
	n ivec3
}

// A general transformation of a sticker can be described by giving a function t
// that transforms vectors.
func (s sticker) transform(t func(ivec3) ivec3) sticker {
	return sticker{
		c: s.c,
		v: t(s.v),
		n: t(s.n),
	}
}

// The following functions return predicates for selecting stickers within a
// given "slice" of a Rubik's cube.

func equalX(x int) func(sticker) bool {
	return func(s sticker) bool {
		return s.v[0] == x
	}
}

func equalY(y int) func(sticker) bool {
	return func(s sticker) bool {
		return s.v[1] == y
	}
}

func equalZ(z int) func(sticker) bool {
	return func(s sticker) bool {
		return s.v[2] == z
	}
}
