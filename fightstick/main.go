package main

import (
	"github.com/deadsy/sdfx/render"
	"github.com/deadsy/sdfx/render/dc"
)

func main() {
	top := topPlane()
	walls := wallsPlane()
	parts := make(Parts)
	parts.add(split2DPlane("top", top, 2))
	parts.add(split2DPlane("walls", walls, 45))

	render.RenderDXF(top, 400, "top.dxf")
	render.RenderDXF(wallsPlane(), 400, "walls.dxf")
	//	render.ToSTL(sdf.Extrude3D(walls, 2), 400, "walls.stl", dc.NewDualContouringDefault())

	for k, v := range parts {
		render.RenderDXF(v.SDF2, 400, k+".dxf")
		render.ToSTL(v.SDF3, 400, k+".stl", dc.NewDualContouringDefault())
	}

}
