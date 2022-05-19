package main

import (
	"github.com/deadsy/sdfx/render"
	"github.com/deadsy/sdfx/sdf"
)

func main() {
	top := topPlane()
	topPlanes := splitPlane()

	for k, v := range topPlanes {
		render.RenderDXF(v, 300, k+".dxf")
		render.RenderSTLSlow(sdf.Extrude3D(v, 2), 1200, k+".stl")
	}
	render.RenderDXF(top, 300, "top.dxf")
}
