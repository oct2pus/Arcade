package main

import (
	"github.com/deadsy/sdfx/render"
	"github.com/deadsy/sdfx/render/dc"
)

func main() {
	render.ToSTL(keycap(), 400, "keycap.stl", dc.NewDualContouringDefault())
	render.ToSTL(walls(), 400, "walls.stl", dc.NewDualContouringDefault())
	render.ToSTL(bottom(), 400, "bottom.stl", dc.NewDualContouringDefault())
	render.ToSTL(top(), 400, "top.stl", dc.NewDualContouringDefault())
	render.ToSTL(plate(), 400, "plate.stl", dc.NewDualContouringDefault())
}
