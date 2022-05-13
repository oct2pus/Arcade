package main

import "github.com/deadsy/sdfx/render"

func main() {
	top := topPlane()
	render.RenderDXF(top, 300, "top.dxf")
	render.RenderSVG(top, 300, "top.svg", "1")
}
