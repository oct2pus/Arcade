package main

import (
	"log"

	"github.com/deadsy/sdfx/sdf"
	v2 "github.com/deadsy/sdfx/vec/v2"
	v3 "github.com/deadsy/sdfx/vec/v3"
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

// split2DPlane splits a 2D plane in half. 0 is right side and 1 is left side.
func split2DPlane(plane sdf.SDF2) []sdf.SDF2 {
	planes := make([]sdf.SDF2, 2)

	rPlane := sdf.Cut2D(plane, v2.Vec{X: 0, Y: 0}, v2.Vec{X: 0, Y: 1})
	planes[0] = rPlane

	lPlane := sdf.Transform2D(sdf.Cut2D(sdf.Transform2D(plane, sdf.MirrorY()), v2.Vec{X: 0, Y: 0}, v2.Vec{X: 0, Y: 1}), sdf.MirrorY())
	planes[1] = lPlane

	return planes
}

// split3DModel splits a 3D model in half. 0 is right side and 1 is left side.
func split3DModel(model sdf.SDF3) []sdf.SDF3 {
	models := make([]sdf.SDF3, 2)
	//plane := sdf.Slice2D(model, v3.Vec{X: 0, Y: 0, Z: 0}, v3.Vec{X: 0, Y: 0, Z: 0})

	//rPlane := sdf.Cut2D(plane, v2.Vec{X: 0, Y: 0}, v2.Vec{X: 0, Y: 1})
	rModel := sdf.Cut3D(model, v3.Vec{X: 0, Y: 0, Z: 0}, v3.Vec{X: 0, Y: 1, Z: 0})
	models[0] = rModel

	//lPlane := sdf.Transform2D(sdf.Cut2D(sdf.Transform2D(plane, sdf.MirrorY()), v2.Vec{X: 0, Y: 0}, v2.Vec{X: 0, Y: 1}), sdf.MirrorY())
	lModel := sdf.Transform3D(sdf.Cut3D(sdf.Transform3D(model, sdf.MirrorXY()), v3.Vec{X: 0, Y: 0, Z: 0}, v3.Vec{X: 0, Y: 1, Z: 0}), sdf.MirrorXY())
	models[1] = lModel

	return models
}
