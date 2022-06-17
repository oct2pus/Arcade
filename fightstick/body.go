package main

import (
	"github.com/deadsy/sdfx/sdf"
	v2 "github.com/deadsy/sdfx/vec/v2"
)

const (
	BODY_SIZE_X    = 300
	BODY_SIZE_Y    = 215
	BODY_SIZE_Z    = 2 + 0 + 0 //Top + Walls + Base
	BODY_CURVE     = 10
	WALL_THICKNESS = 6.9
)

// topPlane produces a 2D top-down image of the fightstick's top panel.
func topPlane() sdf.SDF2 {
	top := sdf.Box2D(v2.Vec{X: BODY_SIZE_X, Y: BODY_SIZE_Y}, BODY_CURVE)

	joystick := joystick(v2.Vec{X: 84, Y: 10 + 10 + 20}) // as listed on jlfmeasure.jpg
	joystick = sdf.Transform2D(joystick, sdf.Rotate2d(sdf.DtoR(90)))
	joystick = loggedMovement(joystick, v2.Vec{X: -top.BoundingBox().Max.X / 2, Y: top.BoundingBox().Max.Y / 7}, "joystick")
	top = sdf.Difference2D(top, joystick)

	buttons := buttonRows()
	buttons = loggedMovement(buttons, v2.Vec{X: top.BoundingBox().Max.X / 2, Y: top.BoundingBox().Max.Y / 4}, "button cluster")
	top = sdf.Difference2D(top, buttons)

	// auxillaryButtons := functionRow()
	// this is a bit ugly
	// auxillaryButtons = loggedMovement(auxillaryButtons, v2.Vec{X: -top.BoundingBox().Max.X / 2.4, Y: 4 * (top.BoundingBox().Max.Y / 5)}, "function cluster")
	// top = sdf.Difference2D(top, auxillaryButtons)

	return top
}

// wallsPlane produces a 2D top-down image of the fightstick's walls.
func wallsPlane() sdf.SDF2 {
	walls := sdf.Box2D(v2.Vec{X: BODY_SIZE_X, Y: BODY_SIZE_Y}, BODY_CURVE)
	body := sdf.Box2D(v2.Vec{X: BODY_SIZE_X - WALL_THICKNESS, Y: BODY_SIZE_Y - WALL_THICKNESS}, BODY_CURVE)
	walls = sdf.Difference2D(walls, body)

	screwProtrosion := sdf.Box2D(v2.Vec{X: WALL_THICKNESS / 3, Y: 12}, 0)
	rProt := loggedMovement(screwProtrosion, v2.Vec{X: (BODY_SIZE_X / 2) - (WALL_THICKNESS / 1.6), Y: 0}, "right screw protrosion")
	lProt := loggedMovement(screwProtrosion, v2.Vec{X: -(BODY_SIZE_X / 2) - -(WALL_THICKNESS / 1.6), Y: 0}, "left screw protrosion")
	tProt := sdf.Transform2D(screwProtrosion, sdf.Rotate2d(sdf.DtoR(90)))
	bProt := loggedMovement(tProt, v2.Vec{X: 0, Y: -(BODY_SIZE_Y / 2) - -(WALL_THICKNESS / 1.6)}, "bottom screw protrosion")
	tProt = loggedMovement(tProt, v2.Vec{X: 0, Y: (BODY_SIZE_Y / 2) - (WALL_THICKNESS / 1.6)}, "top screw protrosion")
	walls = sdf.Union2D(walls, rProt, lProt, tProt, bProt)

	return walls
}
