package main

import (
	"strconv"

	"github.com/deadsy/sdfx/render"
	"github.com/deadsy/sdfx/render/dc"
	"github.com/deadsy/sdfx/sdf"
)

func main() {
	tops := split2DPlane(topPlane())
	bottoms := split2DPlane(bottomPlane())
	render.RenderDXF(topPlane(), 600, "top.dxf")
	render.RenderDXF(wallsPlane(), 600, "walls.dxf")
	render.ToSTL(innerWall(), 400, "innerwall.stl", dc.NewDualContouringDefault())
	render.ToSTL(wallFrontRight(), 400, "wallfrontright.stl", dc.NewDualContouringDefault())
	render.ToSTL(wallBackLeft(), 400, "wallbackleft.stl", dc.NewDualContouringDefault())
	render.ToSTL(wallBackRight(), 400, "wallbackright.stl", dc.NewDualContouringDefault())
	render.ToSTL(wallFrontLeft(), 400, "wallfrontleft.stl", dc.NewDualContouringDefault())
	for i, ele := range tops {
		render.ToSTL(sdf.Extrude3D(ele, 3), 400, "top-"+strconv.Itoa(i)+".stl", dc.NewDualContouringDefault())
	}
	/*for i, ele := range walls {
		render.ToSTL(sdf.Extrude3D(ele, 45), 400, "wall-"+strconv.Itoa(i)+".stl", dc.NewDualContouringDefault())
	}*/
	for i, ele := range bottoms {
		render.ToSTL(sdf.Extrude3D(ele, 3), 400, "bottom-"+strconv.Itoa(i)+".stl", dc.NewDualContouringDefault())
	}
}
