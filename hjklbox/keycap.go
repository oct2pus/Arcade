package main

import (
	"github.com/deadsy/sdfx/sdf"
	v2 "github.com/deadsy/sdfx/vec/v2"
)

// keycap makes a choc switch cap for use with the HJKL and start/select cluster
func keycap() sdf.SDF3 {
	keycapTop2D := sdf.Box2D(sdf.V2{X: 13, Y: 13}, 1)

	keycapBottom2D := sdf.Box2D(sdf.V2{X: 14, Y: 14}, 1)
	keycapBottom2D = sdf.Difference2D(keycapBottom2D, keycapTop2D)

	/*	keycapTop := sdf.Extrude3D(keycapTop2D, 1)
		keycapTop = sdf.Transform3D(keycapTop, sdf.Loft3D())*/
	keycapBottom := sdf.Extrude3D(keycapBottom2D, 1)
	keycapBottom = sdf.Transform3D(keycapBottom, sdf.Translate3d(sdf.V3{X: 0, Y: 0, Z: 1}))
	keycap := sdf.ScaleExtrude3D(keycapTop2D, 1, v2.Vec{X: 1.0775, Y: 1.0775})
	keycap = sdf.Union3D(keycap, keycapBottom)

	stem2D := sdf.Box2D(sdf.V2{X: 1.2, Y: 3}, 0)
	stems := make([]sdf.SDF3, 2)
	for i := range stems {
		stems[i] = sdf.Extrude3D(stem2D, 4)
		stems[i] = sdf.Transform3D(stems[i], sdf.Translate3d(sdf.V3{X: 2.85, Y: 0, Z: 2}))
	}
	stems[1] = sdf.Transform3D(stems[1], sdf.MirrorYZ())

	keycap = sdf.Union3D(keycap, sdf.Union3D(stems...))

	return keycap
}
