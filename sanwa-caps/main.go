package main

import (
	"log"

	"github.com/deadsy/sdfx/render"
	"github.com/deadsy/sdfx/render/dc"
	"github.com/deadsy/sdfx/sdf"
	v2 "github.com/deadsy/sdfx/vec/v2"
	v3 "github.com/deadsy/sdfx/vec/v3"
)

func main() {
	round := 1.0
	topDiameter := 28.0 - round*2
	innerDiameter := 26.5 - round*2
	stemInnerX, stemInnerY := 5.85, 1.55 //0.05 tolerance
	stemOuterDiameter := 7.6
	z := 4.2

	// Top

	top2D, err := sdf.Circle2D(topDiameter / 2)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	// Inside

	inner2D, err := sdf.Circle2D(innerDiameter / 2)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	inner2D = sdf.Difference2D(top2D, inner2D)

	// Stem

	stem2D, err := sdf.Circle2D(stemOuterDiameter / 2)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	stemInner2D := sdf.Box2D(v2.Vec{X: stemInnerX, Y: stemInnerY}, 0)

	stem2D = sdf.Difference2D(stem2D, stemInner2D)

	// Assembly

	//buttCap, err := sdf.ExtrudeRounded3D(top2D, z/4, 0.5)
	buttCap, err := sdf.Loft3D(top2D, inner2D, z, round)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	//body := sdf.Extrude3D(inner2D, z)
	stem := sdf.Extrude3D(stem2D, z)

	buttCap = sdf.Union3D(
		buttCap,
		//		sdf.Transform3D(body, sdf.Translate3d(v3.Vec{X: 0, Y: 0, Z: z/8 + z/2})),
		sdf.Transform3D(stem, sdf.Translate3d(v3.Vec{X: 0, Y: 0, Z: z / 2})),
	)

	render.ToSTL(buttCap, 400, "cap.stl", dc.NewDualContouringDefault())
}
