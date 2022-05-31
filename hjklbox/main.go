package main

import (
	"github.com/deadsy/sdfx/render"
	"github.com/deadsy/sdfx/render/dc"
)

func main() {
	plate := plate()
	top := top()
	bottom := bottom()
	walls := walls()
	render.ToSTL(walls, 400, "walls.stl", dc.NewDualContouringDefault())
	render.ToSTL(bottom, 400, "bottom.stl", dc.NewDualContouringDefault())
	render.ToSTL(top, 400, "top.stl", dc.NewDualContouringDefault())
	render.ToSTL(plate, 400, "plate.stl", dc.NewDualContouringDefault())
}
