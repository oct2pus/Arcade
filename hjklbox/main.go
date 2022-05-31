package main

import (
	"github.com/deadsy/sdfx/render"
	"github.com/deadsy/sdfx/render/dc"
)

func main() {
	//	hjkl := hjkl()
	//	buttons := buttons()
	//	buttonMounts := buttonMounts()
	plate := plate()
	top := top()
	bottom := bottom()
	walls := walls()
	//	holes := usbHoleTest()
	//	render.RenderSTL(holes, 1200, "holes.stl")
	//	render.RenderDXF(hjkl, 300, "hjkl.dxf")
	//	render.RenderDXF(sdf.Difference2D(buttons, buttonMounts), 300, "buttons.dxf")
	//render.RenderSTLSlow(pegholeTest(sdf.V2{X: 47, Y: 11.4}), 1200, "pegholetest.stl")
	//	render.RenderSTLSlow(usbmountHeightTest(), 300, "usbmountHeightTest.stl")
	render.ToSTL(walls, 400, "walls.stl", dc.NewDualContouringDefault())
	render.ToSTL(bottom, 400, "bottom.stl", dc.NewDualContouringDefault())
	render.ToSTL(top, 400, "top.stl", dc.NewDualContouringDefault())
	render.ToSTL(plate, 400, "plate.stl", dc.NewDualContouringDefault())

}
