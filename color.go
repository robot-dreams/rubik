package main

import "fmt"

type color int

const (
	_ color = iota
	white
	yellow
	orange
	red
	green
	blue
)

// These color values came from the excellent post at:
// https://the-rubiks-cube.deviantart.com/journal/Using-Official-Rubik-s-Cube-Colors-268760351
func (c color) rgb() []float32 {
	switch c {
	case white:
		return []float32{1, 1, 1}
	case yellow:
		return []float32{1, 0.835, 0}
	case orange:
		return []float32{1, 0.345, 0}
	case red:
		return []float32{0.769, 0.118, 0.227}
	case green:
		return []float32{0, 0.62, 0.376}
	case blue:
		return []float32{0, 0.318, 0.729}
	default:
		panic(fmt.Sprintf("Invalid color %v", c))
	}
}
