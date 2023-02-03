package main

import (
	"log"

	"github.com/deadsy/sdfx/sdf"
	v2 "github.com/deadsy/sdfx/vec/v2"
)

// Octogon impliments the button interface, which is designed to create an octogonal button.
type Octogon struct {
	size    float64
	rimSize float64
}

// member functions

// Top2D is a topdown view of the top of a keycap.
func (o Octogon) Top2D(size float64) sdf.SDF2 {
	return octogon(size, size)
}

// Rim2D is a topdown view of the sides of a keycap.
func (o Octogon) Rim2D(size float64) sdf.SDF2 {
	rimNegative := octogon(size-o.rimSize, size-o.rimSize)
	return sdf.Difference2D(o.Top2D(size), rimNegative)
}

// Size returns the size value of the Octogon struct.
func (o Octogon) Size() float64 {
	return o.size
}

// new

// octogonNew creates an Octogon struct, which impliments the button interface.
func octogonNew(size, rimSize float64) Octogon {
	var o Octogon
	o.size = size
	o.rimSize = rimSize

	return o
}

// helper functions

// octogon draws an octogon in a 2D plane with a given X and Y value.
func octogon(X, Y float64) sdf.SDF2 {
	dimensions := []v2.Vec{
		{X: -X / 4, Y: -Y / 2},
		{X: -X / 2, Y: -Y / 4},
		{X: -X / 2, Y: Y / 4},
		{X: -X / 4, Y: Y / 2},
		{X: X / 4, Y: Y / 2},
		{X: X / 2, Y: Y / 4},
		{X: X / 2, Y: -Y / 4},
		{X: X / 4, Y: -Y / 2},
	}
	octogon, err := sdf.Polygon2D(dimensions)
	if err != nil {
		log.Fatalln(err)
	}
	return sdf.Center2D(octogon)
}
