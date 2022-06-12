package main

import (
	"log"

	"github.com/deadsy/sdfx/sdf"
	v2 "github.com/deadsy/sdfx/vec/v2"
)

func loggedMovement(input sdf.SDF2, displacement sdf.V2, label string) sdf.SDF2 {
	output := sdf.Transform2D(input, sdf.Translate2d(displacement))
	log.Printf("Moving %v by X: %v, Y: %v\n", label, displacement.X, displacement.Y)
	return output
}

// trapezoid creates a trapezoid from a vector, base is a v2.Vector, xChange modifies the top half.
// positive xChange values make the top half larger, negative xChange values make the top half smaller.
func trapezoid(base v2.Vec, xChange float64) sdf.SDF2 {
	dimensions := v2.VecSet{
		v2.Vec{X: 0, Y: 0},
		v2.Vec{X: -xChange, Y: base.Y},
		v2.Vec{X: base.X + xChange, Y: base.Y},
		v2.Vec{X: base.X, Y: 0},
	}

	trapezoid, err := sdf.Polygon2D(dimensions)
	if err != nil {
		log.Printf("error: %v\nReturning unmodified base.\n", err)
		return sdf.Box2D(base, 0)
	}

	trapezoid = sdf.Center2D(trapezoid)

	return trapezoid
}
