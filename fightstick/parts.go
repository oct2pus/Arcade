package main

import (
	"github.com/deadsy/sdfx/sdf"
	v2 "github.com/deadsy/sdfx/vec/v2"
	v3 "github.com/deadsy/sdfx/vec/v3"
)

type Part struct {
	SDF2 sdf.SDF2
	SDF3 sdf.SDF3
}

type Parts map[string]Part

func (p Parts) add(p2 Parts) {
	for k, v := range p2 {
		p[k] = v
	}
}

func add2DPlane(name string, plane sdf.SDF2, height float64) map[string]Part {
	p := make(map[string]Part)
	p[name] = Part{SDF2: plane, SDF3: sdf.Extrude3D(plane, height)}
	return p
}

func add3DModel(name string, model sdf.SDF3) map[string]Part {
	p := make(map[string]Part)
	p[name] = Part{SDF2: sdf.Slice2D(model, v3.Vec{X: 0, Y: 0, Z: 0}, v3.Vec{X: 0, Y: 0, Z: 0}), SDF3: model}
	return p
}

func split2DPlane(name string, plane sdf.SDF2, height float64) map[string]Part {
	planes := make(map[string]Part)

	rPlane := sdf.Cut2D(plane, v2.Vec{X: 0, Y: 0}, v2.Vec{X: 0, Y: 1})
	planes[name+" "+"right"] = Part{SDF2: rPlane, SDF3: sdf.Extrude3D(rPlane, height)}

	lPlane := sdf.Transform2D(sdf.Cut2D(sdf.Transform2D(plane, sdf.MirrorY()), v2.Vec{X: 0, Y: 0}, v2.Vec{X: 0, Y: 1}), sdf.MirrorY())
	planes[name+" "+"left"] = Part{SDF2: lPlane, SDF3: sdf.Extrude3D(lPlane, height)}

	return planes
}

func split3DModel(name string, model sdf.SDF3) map[string]Part {
	models := make(map[string]Part)
	plane := sdf.Slice2D(model, v3.Vec{X: 0, Y: 0, Z: 0}, v3.Vec{X: 0, Y: 0, Z: 0})

	rPlane := sdf.Cut2D(plane, v2.Vec{X: 0, Y: 0}, v2.Vec{X: 0, Y: 1})
	rModel := sdf.Cut3D(model, v3.Vec{X: 0, Y: 0, Z: 0}, v3.Vec{X: 0, Y: 1, Z: 0})
	models[name+" "+"right"] = Part{SDF2: rPlane, SDF3: rModel}

	lPlane := sdf.Transform2D(sdf.Cut2D(sdf.Transform2D(plane, sdf.MirrorY()), v2.Vec{X: 0, Y: 0}, v2.Vec{X: 0, Y: 1}), sdf.MirrorY())
	lModel := sdf.Transform3D(sdf.Cut3D(sdf.Transform3D(model, sdf.MirrorXY()), v3.Vec{X: 0, Y: 0, Z: 0}, v3.Vec{X: 0, Y: 1, Z: 0}), sdf.MirrorXY())
	models[name+" "+"left"] = Part{SDF2: lPlane, SDF3: lModel}

	return models
}
