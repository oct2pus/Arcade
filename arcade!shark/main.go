package main

import (
	"github.com/deadsy/sdfx/render"
	"github.com/deadsy/sdfx/render/dc"
	"github.com/deadsy/sdfx/sdf"
	v2 "github.com/deadsy/sdfx/vec/v2"
	v3 "github.com/deadsy/sdfx/vec/v3"
)

func main() {
	render.ToSTL(ablzrSwitchHolder(), 200, "ablzrSwitchHolder.stl", dc.NewDualContouringDefault())
	render.ToSTL(ablrzButtonAdapter(), 200, "ablrzButtonAdapter.stl", dc.NewDualContouringDefault())
}

func ablrzButtonAdapter() sdf.SDF3 {
	base2D, _ := sdf.Circle2D(19.6 / 2) // measured 19.8
	// base := sdf.Extrude3D(base2D, 3.5-0.9)
	base := sdf.Extrude3D(base2D, 1)

	line := sdf.Box2D(v2.Vec{X: 1.1, Y: 19.8}, 0)
	line = sdf.Union2D(line, sdf.Transform2D(line, sdf.Rotate2d(sdf.DtoR(90))))

	cross := sdf.Extrude3D(line, 0.5)
	cross = sdf.Transform3D(cross, sdf.Translate3d(v3.Vec{X: 0, Y: 0, Z: base.BoundingBox().Max.Z - cross.BoundingBox().Max.Z}))

	base = sdf.Difference3D(base, cross)
	return base
}

func ablzrSwitchHolder() sdf.SDF3 {
	base2D, _ := triangle(38.5, 34.5, 4)

	choc2D := sdf.Box2D(v2.Vec{X: 13.85, Y: 13.75}, 0)
	choc2D = sdf.Transform2D(choc2D, sdf.Translate2d(v2.Vec{X: 0, Y: -base2D.BoundingBox().Max.Y / 4}))
	base2D = sdf.Difference2D(base2D, choc2D)

	circle3, _ := sdf.Circle2D(3.5 / 2)
	circle45, _ := sdf.Circle2D(4.5 / 2)
	circles := sdf.Union2D(
		sdf.Transform2D(circle45, sdf.Translate2d(v2.Vec{X: 0, Y: base2D.BoundingBox().Max.Y / 1.25})),
		sdf.Transform2D(circle3, sdf.Translate2d(v2.Vec{X: base2D.BoundingBox().Max.X / 1.30, Y: -base2D.BoundingBox().Max.Y / 1.30})),
		sdf.Transform2D(circle3, sdf.Translate2d(v2.Vec{X: -base2D.BoundingBox().Max.X / 1.30, Y: -base2D.BoundingBox().Max.Y / 1.30})),
	)

	base2D = sdf.Difference2D(base2D, circles)
	mount3, _ := sdf.Circle2D(5.5 / 2)
	mount45, _ := sdf.Circle2D(6.5 / 2)
	mounts := sdf.Union2D(
		sdf.Transform2D(mount45, sdf.Translate2d(v2.Vec{X: 0, Y: base2D.BoundingBox().Max.Y / 1.25})),
		sdf.Transform2D(mount3, sdf.Translate2d(v2.Vec{X: base2D.BoundingBox().Max.X / 1.30, Y: -base2D.BoundingBox().Max.Y / 1.30})),
		sdf.Transform2D(mount3, sdf.Translate2d(v2.Vec{X: -base2D.BoundingBox().Max.X / 1.30, Y: -base2D.BoundingBox().Max.Y / 1.30})),
	)

	mounts = sdf.Difference2D(mounts, circles)

	pegs := sdf.Extrude3D(mounts, 3)
	base := sdf.Extrude3D(base2D, 1.6)

	pegs = sdf.Transform3D(pegs, sdf.Translate3d(v3.Vec{X: 0, Y: 0, Z: -base.BoundingBox().Max.Z + pegs.BoundingBox().Max.Z}))

	return sdf.Union3D(base, pegs)
}

/*func triangle(sides float64) (sdf.SDF2, error) {
	dimensions := []v2.Vec{{X: -sides / 2, Y: 0}, {X: 0, Y: sides / 2}, {X: sides / 2, Y: 0}}
	return sdf.Polygon2D(dimensions)
}*/

func triangle(base, height, trim float64) (sdf.SDF2, error) {
	dimensions := []v2.Vec{
		{X: (-base / 2) - (-trim), Y: (-height / 2)},
		{X: (-base / 2) - (-trim / 2), Y: (-height / 2) + (trim)},
		{X: (-trim / 2), Y: (height / 2) - (trim)},
		{X: (trim / 2), Y: (height / 2) - (trim)},
		{X: (base / 2) - (trim / 2), Y: (-height / 2) + trim},
		{X: (base / 2) - (trim), Y: (-height / 2)},
	}
	triangle, err := sdf.Polygon2D(dimensions)
	// this might cause issues lol
	triangle = sdf.Center2D(triangle)
	return triangle, err
}
