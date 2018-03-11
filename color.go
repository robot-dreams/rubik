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

func (c color) rgb() []float32 {
	switch c {
	case white:
		return []float32{1, 1, 1}
	case yellow:
		return []float32{0.9, 0.9, 0}
	case orange:
		return []float32{0.9, 0.4, 0.1}
	case red:
		return []float32{0.8, 0, 0}
	case green:
		return []float32{0, 0.7, 0.2}
	case blue:
		return []float32{0, 0.2, 0.9}
	default:
		panic(fmt.Sprintf("Unknown color %v", c))
	}
}
