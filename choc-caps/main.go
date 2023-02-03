package main

import (
	"github.com/deadsy/sdfx/render"
	"github.com/deadsy/sdfx/render/dc"
)

const (
	// these values are designed to fit into the default flatbox,
	// tinker with the values to find the size you need.
	OCTOGON_SIZE_SMALL = 20.25
	OCTOGON_SIZE_LARGE = 25.75
	CIRCLE_SIZE_SMALL  = 22.5
	CIRCLE_SIZE_LARGE  = 28.5
	// A larger value equals a wider rim, this is subtractive.
	RIM_SIZE = 1.0
	// These control how tall each individual element is. Keep in mind these values are absolute,
	// not relative. For example, the default MID_Z extends 2.0mm past the TOP_Z.
	TOP_Z  = 1.5
	MID_Z  = 3.5
	STEM_Z = 5
	// ROUND determines the radius of curves, 0 is sharp edges.
	// keep that in mind if you're modifying this program.
	// each individual Z height *must* be greater than round * 2.
	ROUND = 0.5
)

// Larger values in render.ToSTL will produce higher quality models, but also larger filesizes
// the larger you make it, the more computationally intensive it is.
// this model uses dc.NewDualContouringDefault
// this renders very slowly and is more computationally intensive
// than two other renderers. i find it has the best visual clarity when 300 or above
// which produces a smaller filesize that I can actually share online
// the other two renderers are render.NewMarchingCubesOctree and NewMarchingCubesUniform
// they compile faster but need significantly higher values for acceptable clarity,
// which in turn creates significantly larger file sizes.

func main() {
	oS, oL := octogonNew(OCTOGON_SIZE_SMALL, RIM_SIZE), octogonNew(OCTOGON_SIZE_LARGE, RIM_SIZE)
	cS, cL := circleNew(CIRCLE_SIZE_SMALL, RIM_SIZE), circleNew(CIRCLE_SIZE_LARGE, RIM_SIZE)
	render.ToSTL(create(cS), "circularSmall.stl", dc.NewDualContouringDefault(300))
	render.ToSTL(create(cL), "circularLarge.stl", dc.NewDualContouringDefault(300))
	render.ToSTL(create(oS), "octogonalSmall.stl", dc.NewDualContouringDefault(300))
	render.ToSTL(create(oL), "octogonalLarge.stl", dc.NewDualContouringDefault(300))
}
