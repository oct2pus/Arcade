package main

import (
	"github.com/deadsy/sdfx/render"
)

func main() {
	//	hjkl := hjkl()
	//	buttons := buttons()
	//	buttonMounts := buttonMounts()
	plate := plate()
	top := top()
	bottom := bottom()
	//	holes := usbHoleTest()
	//	render.RenderSTL(holes, 1200, "holes.stl")
	//	render.RenderDXF(hjkl, 300, "hjkl.dxf")
	//	render.RenderDXF(sdf.Difference2D(buttons, buttonMounts), 300, "buttons.dxf")
	render.RenderSTLSlow(bottom, 1200, "bottom.stl")
	render.RenderSTLSlow(top, 1200, "top.stl")
	render.RenderSTLSlow(plate, 1200, "plate.stl")

}
