package main

import (
	"github.com/deadsy/sdfx/obj"
	"github.com/deadsy/sdfx/render"
	"github.com/deadsy/sdfx/sdf"
	v3 "github.com/deadsy/sdfx/vec/v3"
)

func main() {
	plug24mm := plug(26, 23, 25)
	render.ToSTL(plug24mm, "24mm_plug.stl", render.NewMarchingCubesUniform(300))
	render.ToSTL(nut(32, 6, plug24mm), "24mm_nut.stl", render.NewMarchingCubesUniform(300))
}

func plug(capDimeter, plugDiameter, length float64) sdf.SDF3 {
	thread, _ := sdf.ISOThread(plugDiameter/2, 2, true)
	plug, _ := sdf.Screw3D(thread, length, 0, 2, 1)

	cRad1, _ := sdf.Circle2D((capDimeter - 4) / 2)
	cRad2, _ := sdf.Circle2D(capDimeter / 2)

	cap, _ := sdf.Loft3D(cRad1, cRad2, 3.75, 1)
	cap = sdf.Transform3D(cap, sdf.RotateX(sdf.DtoR(180)))
	cap = sdf.Transform3D(cap, sdf.Translate3d(v3.Vec{X: 0, Y: 0, Z: -length/2 - (3.75)/2}))
	return sdf.Union3D(cap, plug)
}

func nut(headDiameter, height float64, plug sdf.SDF3) sdf.SDF3 {
	nut, _ := obj.HexHead3D(headDiameter/2, height, "tb")
	nut = sdf.Difference3D(nut, plug)

	nut = sdf.Transform3D(nut, sdf.Scale3d(v3.Vec{X: 1.05, Y: 1.05, Z: 1}))
	return nut
}
