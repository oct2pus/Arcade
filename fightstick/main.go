package main

import (
	"github.com/deadsy/sdfx/render"
	"github.com/deadsy/sdfx/render/dc"
	"github.com/deadsy/sdfx/sdf"
)

func main() {
	top := topPlane()
	walls := wallsPlane()
	planes := make(planes)
	planes.add(splitPlane("top", top))
	planes.add(splitPlane("walls", walls))

	render.RenderDXF(top, 400, "top.dxf")
	render.RenderDXF(wallsPlane(), 400, "walls.dxf")
	//	render.ToSTL(sdf.Extrude3D(walls, 2), 400, "walls.stl", dc.NewDualContouringDefault())

	for k, v := range planes {
		render.RenderDXF(v, 400, k+".dxf")
		render.ToSTL(sdf.Extrude3D(v, 2), 400, k+".stl", dc.NewDualContouringDefault())
	}

}
