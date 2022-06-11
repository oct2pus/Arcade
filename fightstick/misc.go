package main

import (
	"log"

	"github.com/deadsy/sdfx/sdf"
)

func loggedMovement(input sdf.SDF2, displacement sdf.V2, label string) sdf.SDF2 {
	output := sdf.Transform2D(input, sdf.Translate2d(displacement))
	log.Printf("Moving %v by X: %v, Y: %v\n", label, displacement.X, displacement.Y)
	return output
}

func splitPlane() map[string]sdf.SDF2 {
	planes := make(map[string]sdf.SDF2)
	original := topPlane()
	x, y := original.BoundingBox().Max.X*2, original.BoundingBox().Max.Y*2
	cutout := sdf.Box2D(sdf.V2{X: x, Y: y}, 0)
	cOLeft := sdf.Transform2D(cutout, sdf.Translate2d(sdf.V2{X: x / 2, Y: 0}))
	cORight := sdf.Transform2D(cutout, sdf.Translate2d(sdf.V2{X: -x / 2, Y: 0}))
	planes["left"] = sdf.Difference2D(original, cOLeft)
	planes["right"] = sdf.Difference2D(original, cORight)
	return planes
}
