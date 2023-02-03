package main

import (
	"log"

	"github.com/deadsy/sdfx/sdf"
	v2 "github.com/deadsy/sdfx/vec/v2"
	v3 "github.com/deadsy/sdfx/vec/v3"
)

// button interface impliments the 3 seperate sections of a choc keycap
type button interface {
	Rim2D(size float64) sdf.SDF2
	Top2D(size float64) sdf.SDF2
	Size() float64
}

// 3D

// create combines all the 2d parts of a choc buttons and creates a 3d model of the button
func create(b button) sdf.SDF3 {
	top, err := sdf.ExtrudeRounded3D(b.Top2D(b.Size()-ROUND*2), TOP_Z, ROUND)
	if err != nil {
		log.Fatalln(err)
	}
	mid, err := sdf.ExtrudeRounded3D(b.Rim2D(b.Size()-ROUND*2), MID_Z, ROUND)
	if err != nil {
		log.Fatalln(err)
	}
	stem := stem()

	button := sdf.Union3D(
		sdf.Transform3D(top, sdf.Translate3d(v3.Vec{X: 0, Y: 0, Z: -STEM_Z/2 + TOP_Z/2})),
		sdf.Transform3D(mid, sdf.Translate3d(v3.Vec{X: 0, Y: 0, Z: -STEM_Z/2 + MID_Z/2})),
		sdf.Transform3D(stem, sdf.Translate3d(v3.Vec{X: 0, Y: 0, Z: TOP_Z / 3})),
	)

	//button = sdf.Transform3D(button, sdf.RotateX(sdf.DtoR(180)))
	return button
}

// stem is the 2d stem of the button
func stem() sdf.SDF3 {
	stem2D := stem2D()
	stem := sdf.Extrude3D(stem2D, STEM_Z)
	// this was intended to make it easier to print vertically, but i later gave up on the idea.
	// you can use this to narrow the stem towards the end so that its a smaller surface area that is easier to remove from supports
	/*	cutoutDiameter := 1.0
		cutout2D, err := sdf.Circle2D(cutoutDiameter / 2)
		if err != nil {
			log.Fatalln(err)
		}
		cutout2D = sdf.Union2D(
			sdf.Transform2D(cutout2D, sdf.Translate2d(v2.Vec{X: stem2D.BoundingBox().Size().Y / 2, Y: -stem2D.BoundingBox().Size().Y - (-cutoutDiameter / 4)})),
			sdf.Transform2D(cutout2D, sdf.Translate2d(v2.Vec{X: -stem2D.BoundingBox().Size().Y / 2, Y: -stem2D.BoundingBox().Size().Y - (-cutoutDiameter / 4)})),
		)

		cutout := sdf.Extrude3D(cutout2D, stem2D.BoundingBox().Size().X*2)
		cutout = sdf.Transform3D(cutout, sdf.RotateY(sdf.DtoR(90)))
		cutout = sdf.Transform3D(cutout, sdf.RotateX(sdf.DtoR(90)))

		cutout = sdf.Union3D(
			sdf.Transform3D(cutout, sdf.Translate3d(v3.Vec{X: 0, Y: stem2D.BoundingBox().Size().X / 2, Z: 0})),
			sdf.Transform3D(cutout, sdf.Translate3d(v3.Vec{X: 0, Y: -stem2D.BoundingBox().Size().X / 2, Z: 0})),
		)
		stem = sdf.Difference3D(stem, cutout) */
	return stem
}

//2D

// stem2D is a top down view  of the keycap's stem, it has cutouts to make it easier to remove/insert the keycap from the stem.
// This is supposed to make it less likely you'll break the stem removing it.
func stem2D() sdf.SDF2 {
	x, y := 1.1, 2.8
	spacing := 5.8
	cutout := sdf.Box2D(v2.Vec{X: x / 1.8, Y: y / 1.5}, 0)
	stem := sdf.Box2D(v2.Vec{X: x, Y: y}, 0)
	cutout = sdf.Union2D(
		sdf.Transform2D(cutout, sdf.Translate2d(v2.Vec{X: -x / 2, Y: 0})),
		sdf.Transform2D(cutout, sdf.Translate2d(v2.Vec{X: x / 2, Y: 0})),
	)
	stem = sdf.Difference2D(stem, cutout)
	stems := sdf.Union2D(
		sdf.Transform2D(stem, sdf.Translate2d(v2.Vec{X: -spacing / 2, Y: 0})),
		sdf.Transform2D(stem, sdf.Translate2d(v2.Vec{X: spacing / 2, Y: 0})),
	)
	return sdf.Center2D(stems)
}
