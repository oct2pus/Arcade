package main

import (
	"github.com/deadsy/sdfx/render"
)

func main() {
	insert, _ := Insert()
	balltop, _ := balltop()
	dustCover, _ := dustCover()
	shaftCover, _ := shaftCover()
	render.RenderSTL(insert, 300, "insert.stl")
	render.RenderSTL(balltop, 300, "balltop.stl")
	render.RenderSTL(dustCover, 300, "dustcover.stl")
	render.RenderSTL(shaftCover, 300, "shaftcover.stl")
}
