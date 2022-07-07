package main

import (
	"github.com/deadsy/sdfx/sdf"
	v2 "github.com/deadsy/sdfx/vec/v2"
	v3 "github.com/deadsy/sdfx/vec/v3"
)

// test button diameter and neutrik holes
func holeTest() sdf.SDF3 {
	body := sdf.Box2D(v2.Vec{X: 60, Y: 40}, 0)
	body = sdf.Difference2D(body, sdf.Transform2D(neutrik(), sdf.Translate2d(v2.Vec{X: -16, Y: 0})))
	button, _ := sdf.Circle2D(BUTTON30_DIAMETER / 2)
	body = sdf.Difference2D(body, sdf.Transform2D(button, sdf.Translate2d(v2.Vec{X: 13, Y: 0})))

	return sdf.Extrude3D(body, 2)
}

// 7.2 is a good number for M4 screws
// 5.6 is a good number for M3 Screws
func countersinkTest() sdf.SDF3 {
	base2D := sdf.Box2D(v2.Vec{X: 10, Y: 10}, 0)
	hole2D, _ := sdf.Circle2D(M3_SCREW_DIAMETER / 2)
	base2D = sdf.Difference2D(base2D, hole2D)

	base := sdf.Extrude3D(base2D, 5)
	// cone 3d rapidly balloons the filesize lol
	cone, _ := sdf.Cone3D(2.3, M3_SCREW_DIAMETER/2, 5.6/2, 0)
	/*	indent2D, _ := sdf.Circle2D(9.2 / 2)
		indent := sdf.Extrude3D(indent2D, 2.3)
		indent = sdf.Transform3D(indent, sdf.Translate3d(v3.Vec{X: 0, Y: 0, Z: base.BoundingBox().Max.Z}))

		base = sdf.Difference3D(base, indent)*/
	cone = sdf.Transform3D(cone, sdf.Translate3d(v3.Vec{X: 0, Y: 0, Z: base.BoundingBox().Max.Z - cone.BoundingBox().Max.Z}))
	base = sdf.Difference3D(base, cone)
	return base
}
