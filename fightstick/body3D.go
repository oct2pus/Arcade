package main

import (
	"github.com/deadsy/sdfx/sdf"
	v2 "github.com/deadsy/sdfx/vec/v2"
	v3 "github.com/deadsy/sdfx/vec/v3"
)

const (
	TOP_HEIGHT    = 4.0
	WALLS_HEIGHT  = 45.0
	BOTTOM_HEIGHT = 3.0
	WALL_NOTCH    = 5.0
)

func top() sdf.SDF3 {
	top := sdf.Extrude3D(topPlane(), TOP_HEIGHT)
	//	screws := sdf.Transform3D(screwCountersinks(), sdf.Translate3d(v3.Vec{X: 0, Y: 0, Z: TOP_HEIGHT/2 - screwCountersinks().BoundingBox().Max.Z}))
	//	top = sdf.Difference3D(top, screws)
	return top
}

func topLeft() sdf.SDF3 {
	return split3DModel(top())[0]
}

func topRight() sdf.SDF3 {
	return split3DModel(top())[1]

}

// wallFrontRight is the front right wall. This houses the neutrik connector.
func wallFrontRight() sdf.SDF3 {
	corner := wallCorner()

	neutrik := sdf.Extrude3D(neutrik(), WALL_THICKNESS)
	neutrik = sdf.Transform3D(neutrik, sdf.RotateZ(sdf.DtoR(90)))
	neutrik = sdf.Transform3D(neutrik, sdf.RotateY(sdf.DtoR(90)))
	neutrik = sdf.Transform3D(neutrik, sdf.MirrorXY())
	neutrik = sdf.Transform3D(neutrik, sdf.Translate3d(v3.Vec{X: BODY_SIZE_X/3 + (WALL_THICKNESS / 2), Y: BODY_SIZE_Y / 3, Z: 0}))
	corner = sdf.Difference3D(corner, neutrik)

	corner = sdf.Transform3D(corner, sdf.MirrorYZ())
	corner = sdf.Transform3D(corner, sdf.Rotate3d(v3.Vec{X: 0, Y: 0, Z: 1}, sdf.DtoR(270)))

	return corner
}

//wallFrontLeft is the front left wall. This houses 4 24mm buttons.
func wallFrontLeft() sdf.SDF3 {
	corner := wallCorner()
	// Could be removed, just need to modify how functionButtons move
	corner = sdf.Transform3D(corner, sdf.Rotate3d(v3.Vec{X: 0, Y: 0, Z: 1}, sdf.DtoR(270)))

	functionButtons := sdf.Extrude3D(functionRow(), WALL_THICKNESS)
	functionButtons = sdf.Transform3D(functionButtons, sdf.RotateX(sdf.DtoR(90)))
	functionButtons = sdf.Transform3D(functionButtons, sdf.Translate3d(v3.Vec{X: BODY_SIZE_X/4.5 + (WALL_THICKNESS / 2), Y: -BODY_SIZE_Y / 2, Z: 0}))
	corner = sdf.Difference3D(corner, functionButtons)

	corner = sdf.Transform3D(corner, sdf.Rotate3d(v3.Vec{X: 0, Y: 0, Z: 1}, sdf.DtoR(180)))

	return corner
}

//wallBackRight is the back right wall.
func wallBackRight() sdf.SDF3 {
	corner := wallCorner()
	corner = sdf.Transform3D(corner, sdf.RotateZ(sdf.DtoR(270)))
	return corner
}

//wallBackLeft is the back left wall.
func wallBackLeft() sdf.SDF3 {
	corner := wallCorner()
	corner = sdf.Transform3D(corner, sdf.RotateZ(sdf.DtoR(270)))
	corner = sdf.Transform3D(corner, sdf.MirrorYZ())
	return corner
}

// inner wall is an internal wall to hold the fightstick together.
func innerWall() sdf.SDF3 {
	wall := sdf.Extrude3D(innerWallPlane(), WALLS_HEIGHT)

	// cut off edges
	end, _ := sdf.Box3D(v3.Vec{X: INNER_WALL_WIDTH, Y: WALL_THICKNESS, Z: WALLS_HEIGHT}, 0)
	cutout, _ := sdf.Box3D(v3.Vec{X: 12, Y: INNER_WALL_WIDTH, Z: WALLS_HEIGHT - (WALL_NOTCH+0.3)*2}, 0)
	endCutout := sdf.Difference3D(end, cutout)

	wall = sdf.Difference3D(wall, sdf.Transform3D(endCutout, sdf.Translate3d(v3.Vec{X: 0, Y: BODY_SIZE_Y/2 - WALL_THICKNESS/2, Z: 0})))
	wall = sdf.Difference3D(wall, sdf.Transform3D(endCutout, sdf.Translate3d(v3.Vec{X: 0, Y: -BODY_SIZE_Y/2 - -WALL_THICKNESS/2, Z: 0})))
	//fill the center holes

	filler, _ := sdf.Box3D(v3.Vec{X: INNER_WALL_WIDTH, Y: WALL_THICKNESS, Z: 10}, 0) // Z is arbitrary, bigger than center portion
	wall = sdf.Union3D(wall, filler)

	// create holes for wires to pass through
	centerCutout := sdf.Extrude3D(trapezoid(v2.Vec{X: BODY_SIZE_Y / 3, Y: WALLS_HEIGHT / 3}, -WALLS_HEIGHT/3), INNER_WALL_WIDTH)
	centerCutout = sdf.Transform3D(centerCutout, sdf.RotateX(sdf.DtoR(90)))
	centerCutout = sdf.Transform3D(centerCutout, sdf.RotateZ(sdf.DtoR(90)))

	rotatedCenterCutout := sdf.Transform3D(centerCutout, sdf.RotateX(sdf.DtoR(180)))

	// here comes the ugly bit
	wall = sdf.Difference3D(wall, sdf.Transform3D(rotatedCenterCutout, sdf.Translate3d(v3.Vec{X: 0, Y: BODY_SIZE_Y/2 - centerCutout.BoundingBox().Max.Y*1.3, Z: -WALLS_HEIGHT/2 - -centerCutout.BoundingBox().Max.Z*1.5})))
	wall = sdf.Difference3D(wall, sdf.Transform3D(centerCutout, sdf.Translate3d(v3.Vec{X: 0, Y: 0, Z: -WALLS_HEIGHT/2 - -centerCutout.BoundingBox().Max.Z*1.5})))
	wall = sdf.Difference3D(wall, sdf.Transform3D(rotatedCenterCutout, sdf.Translate3d(v3.Vec{X: 0, Y: -BODY_SIZE_Y/2 - -centerCutout.BoundingBox().Max.Y*1.3, Z: -WALLS_HEIGHT/2 - -centerCutout.BoundingBox().Max.Z*1.5})))
	// FLIP IT TURNWAYS
	wall = sdf.Difference3D(wall, sdf.Transform3D(centerCutout, sdf.Translate3d(v3.Vec{X: 0, Y: BODY_SIZE_Y/2 - centerCutout.BoundingBox().Max.Y*1.3, Z: WALLS_HEIGHT/2 - centerCutout.BoundingBox().Max.Z*1.5})))
	wall = sdf.Difference3D(wall, sdf.Transform3D(rotatedCenterCutout, sdf.Translate3d(v3.Vec{X: 0, Y: 0, Z: WALLS_HEIGHT/2 - centerCutout.BoundingBox().Max.Z*1.5})))
	wall = sdf.Difference3D(wall, sdf.Transform3D(centerCutout, sdf.Translate3d(v3.Vec{X: 0, Y: -BODY_SIZE_Y/2 - -centerCutout.BoundingBox().Max.Y*1.3, Z: WALLS_HEIGHT/2 - centerCutout.BoundingBox().Max.Z*1.5})))

	return wall
}

func wallCorner() sdf.SDF3 {
	corner := sdf.Extrude3D(wallCornerPlane(), WALLS_HEIGHT)
	cutout, _ := sdf.Box3D(v3.Vec{X: 12, Y: INNER_WALL_WIDTH, Z: WALLS_HEIGHT - WALL_NOTCH*2}, 0) //X is arbitrary

	corner = sdf.Difference3D(corner, sdf.Transform3D(cutout, sdf.Translate3d(v3.Vec{X: BODY_SIZE_X/3 + cutout.BoundingBox().Max.X, Y: 0, Z: 0})))
	return corner
}
