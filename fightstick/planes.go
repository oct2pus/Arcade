package main

import (
	"github.com/deadsy/sdfx/sdf"
	v2 "github.com/deadsy/sdfx/vec/v2"
)

type planes map[string]sdf.SDF2

func (p planes) add(p2 planes) {
	for k, v := range p2 {
		p[k] = v
	}
}

func splitPlane(name string, plane sdf.SDF2) map[string]sdf.SDF2 {
	planes := make(map[string]sdf.SDF2)
	/*	original := topPlane()
		x, y := original.BoundingBox().Max.X*2, original.BoundingBox().Max.Y*2
		cutout := sdf.Box2D(sdf.V2{X: x, Y: y}, 0)
		cOLeft := sdf.Transform2D(cutout, sdf.Translate2d(sdf.V2{X: x / 2, Y: 0}))
		cORight := sdf.Transform2D(cutout, sdf.Translate2d(sdf.V2{X: -x / 2, Y: 0}))
		planes["left"] = sdf.Difference2D(original, cOLeft)
		planes["right"] = sdf.Difference2D(original, cORight)*/
	planes[name+" "+"right"] = sdf.Cut2D(plane, v2.Vec{X: 0, Y: 0}, v2.Vec{X: 0, Y: 1})
	planes[name+" "+"left"] = sdf.Transform2D(sdf.Cut2D(sdf.Transform2D(plane, sdf.MirrorY()), v2.Vec{X: 0, Y: 0}, v2.Vec{X: 0, Y: 1}), sdf.MirrorY())
	return planes
}
