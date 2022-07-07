package main

import (
	"github.com/deadsy/sdfx/sdf"
	v2 "github.com/deadsy/sdfx/vec/v2"
)

const (
	BODY_SIZE_X      = 300
	BODY_SIZE_Y      = 215
	BODY_SIZE_Z      = 3 + 45 + 3 //Top + Walls + Base
	BODY_CURVE       = 10
	WALL_THICKNESS   = 6.9
	INNER_WALL_WIDTH = 12.0
)

// topPlane produces a 2D top-down image of the fightstick's top panel.
// TODO: need to add countersinks for the screws lol.
func topPlane() sdf.SDF2 {
	top := sdf.Box2D(v2.Vec{X: BODY_SIZE_X, Y: BODY_SIZE_Y}, BODY_CURVE)

	joystick := joystick(v2.Vec{X: 84, Y: 10 + 10 + 20}) // as listed on jlfmeasure.jpg
	joystick = sdf.Transform2D(joystick, sdf.Rotate2d(sdf.DtoR(90)))
	joystick = loggedMovement(joystick, v2.Vec{X: -top.BoundingBox().Max.X / 2, Y: top.BoundingBox().Max.Y / 7}, "joystick")
	top = sdf.Difference2D(top, joystick)

	buttons := buttonRows()
	buttons = loggedMovement(buttons, v2.Vec{X: top.BoundingBox().Max.X / 2, Y: top.BoundingBox().Max.Y / 4}, "button cluster")
	top = sdf.Difference2D(top, buttons)

	top = sdf.Difference2D(top, screwHoles())

	return top
}

// wallsPlane produces a 2D top-down image of the fightstick's walls.
func wallsPlane() sdf.SDF2 {
	walls := sdf.Box2D(v2.Vec{X: BODY_SIZE_X, Y: BODY_SIZE_Y}, BODY_CURVE)
	body := sdf.Box2D(v2.Vec{X: BODY_SIZE_X - WALL_THICKNESS, Y: BODY_SIZE_Y - WALL_THICKNESS}, BODY_CURVE)
	walls = sdf.Difference2D(walls, body)

	screwHole := sdf.Box2D(v2.Vec{X: WALL_THICKNESS / 3, Y: WALL_THICKNESS * 2.2}, 0)
	rHole := loggedMovement(screwHole, v2.Vec{X: (BODY_SIZE_X / 2) - (WALL_THICKNESS / 1.6), Y: 0}, "right screw hole")
	lHole := loggedMovement(screwHole, v2.Vec{X: -(BODY_SIZE_X / 2) - -(WALL_THICKNESS / 1.6), Y: 0}, "left screw hole")
	tHole := sdf.Transform2D(screwHole, sdf.Rotate2d(sdf.DtoR(90)))
	bHole := loggedMovement(tHole, v2.Vec{X: 0, Y: -(BODY_SIZE_Y / 2) - -(WALL_THICKNESS / 1.6)}, "bottom screw hole")
	tHole = loggedMovement(tHole, v2.Vec{X: 0, Y: (BODY_SIZE_Y / 2) - (WALL_THICKNESS / 1.6)}, "top screw hole")
	walls = sdf.Union2D(walls, rHole, lHole, tHole, bHole)

	corner := trapezoid(v2.Vec{X: WALL_THICKNESS, Y: WALL_THICKNESS}, WALL_THICKNESS)
	brCorner := sdf.Transform2D(corner, sdf.Rotate2d(sdf.DtoR(45)))
	brCorner = loggedMovement(brCorner, v2.Vec{X: (BODY_SIZE_X / 2) - (WALL_THICKNESS), Y: -(BODY_SIZE_Y / 2) - -(WALL_THICKNESS)}, "bottom right wall corner")
	blCorner := sdf.Transform2D(brCorner, sdf.MirrorY())
	trCorner := sdf.Transform2D(brCorner, sdf.MirrorX())
	tlCorner := sdf.Transform2D(trCorner, sdf.MirrorY())
	walls = sdf.Union2D(walls, brCorner, blCorner, trCorner, tlCorner)

	walls = sdf.Difference2D(walls, screwHoles())

	return walls
}

func bottomPlane() sdf.SDF2 {
	bottom := sdf.Box2D(v2.Vec{X: BODY_SIZE_X, Y: BODY_SIZE_Y}, BODY_CURVE)

	bottom = sdf.Difference2D(bottom, screwHoles())

	return bottom
}

func innerWallPlane() sdf.SDF2 {
	wall := sdf.Box2D(v2.Vec{X: INNER_WALL_WIDTH, Y: BODY_SIZE_Y}, 0)

	wall = sdf.Difference2D(wall, screwHoles())

	return wall
}

// wallCorner returns one corner of the wall.
func wallCornerPlane() sdf.SDF2 {
	segmentPlane := wallsPlane()
	segmentPlane = split2DPlane(segmentPlane)[0]
	segmentPlane = sdf.Center2D(segmentPlane)
	segmentPlane = sdf.Transform2D(segmentPlane, sdf.Rotate2d(sdf.DtoR(90)))
	segmentPlane = sdf.Center2D(segmentPlane)
	segmentPlane = split2DPlane(segmentPlane)[0]
	segmentPlane = sdf.Center2D(segmentPlane)
	return segmentPlane
}
