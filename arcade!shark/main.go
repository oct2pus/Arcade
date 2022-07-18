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
	base := sdf.Extrude3D(base2D, 3.5-0.9)
	line := sdf.Box2D(v2.Vec{X: 1.1, Y: 19.8}, 0)
	line = sdf.Union2D(line, sdf.Transform2D(line, sdf.Rotate2d(sdf.DtoR(90))))
	cross := sdf.Extrude3D(line, 1)
	cross = sdf.Transform3D(cross, sdf.Translate3d(v3.Vec{X: 0, Y: 0, Z: base.BoundingBox().Max.Z - cross.BoundingBox().Max.Z}))
	base = sdf.Difference3D(base, cross)
	return base
}

func ablzrSwitchHolder() sdf.SDF3 {
	base2D, _ := triangle(38, 34, 4)

	return sdf.Extrude3D(base2D, 2)
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

	return sdf.Polygon2D(dimensions)
}
