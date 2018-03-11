package main

// There are 6 faces on a Rubik's Cube, and each face has 3x3 stickers.
type rubiksCube [6 * 3 * 3]sticker

func newRubiksCube() rubiksCube {
	var r rubiksCube
	stickers := make([]sticker, 0, len(r))
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			stickers = append(stickers, sticker{white, ivec3{i, 1, j}, posY})
			stickers = append(stickers, sticker{yellow, ivec3{i, -1, j}, negY})
			stickers = append(stickers, sticker{orange, ivec3{1, i, j}, posX})
			stickers = append(stickers, sticker{red, ivec3{-1, i, j}, negX})
			stickers = append(stickers, sticker{blue, ivec3{i, j, 1}, posZ})
			stickers = append(stickers, sticker{green, ivec3{i, j, -1}, negZ})
		}
	}
	copy(r[:], stickers)
	return r
}

// A general transformation of a Rubik's Cube can be described by giving a
// predicate p that specifies which stickers to transform, together with a
// function t that transforms vectors.
func (r *rubiksCube) transform(p func(sticker) bool, t func(ivec3) ivec3) {
	for i, s := range *r {
		if p(s) {
			(*r)[i] = s.transform(t)
		}
	}
}

// The following methods rotate a "slice" of a Rubik's Cube 90 degrees
// counter-clockwise about a given axis.

func (r *rubiksCube) rotateX(x int) {
	r.transform(equalX(x), rotateX)
}

func (r *rubiksCube) rotateY(y int) {
	r.transform(equalY(y), rotateY)
}

func (r *rubiksCube) rotateZ(z int) {
	r.transform(equalZ(z), rotateZ)
}
