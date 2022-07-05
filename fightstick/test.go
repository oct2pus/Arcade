package main

import (
	"github.com/deadsy/sdfx/sdf"
	v2 "github.com/deadsy/sdfx/vec/v2"
)

// test button diameter and neutrik holes
func holeTest() sdf.SDF3 {
	body := sdf.Box2D(v2.Vec{X: 60, Y: 40}, 0)
	body = sdf.Difference2D(body, sdf.Transform2D(neutrik(), sdf.Translate2d(v2.Vec{X: -16, Y: 0})))
	button, _ := sdf.Circle2D(BUTTON30_DIAMETER / 2)
	body = sdf.Difference2D(body, sdf.Transform2D(button, sdf.Translate2d(v2.Vec{X: 13, Y: 0})))

	return sdf.Extrude3D(body, 2)
}
