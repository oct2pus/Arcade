package main

import (
	"github.com/deadsy/sdfx/sdf"
	v2 "github.com/deadsy/sdfx/vec/v2"
)

const (
	BODY_SIZE_X    = 300
	BODY_SIZE_Y    = 220
	BODY_SIZE_Z    = 2 + 0 + 0 //Top + Walls + Base
	BODY_CURVE     = 10
	WALL_THICKNESS = 15
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

	auxillaryButtons := functionRow()
	// this is a bit ugly
	auxillaryButtons = loggedMovement(auxillaryButtons, v2.Vec{X: -top.BoundingBox().Max.X / 2.4, Y: 4 * (top.BoundingBox().Max.Y / 5)}, "function cluster")
	top = sdf.Difference2D(top, auxillaryButtons)

	return top
}

// wallsPlane produces a 2D top-down image of the fightstick's walls.
func wallsPlane() sdf.SDF2 {
	walls := sdf.Box2D(v2.Vec{X: BODY_SIZE_X, Y: BODY_SIZE_Y}, BODY_CURVE)
	body := sdf.Box2D(v2.Vec{X: BODY_SIZE_X - WALL_THICKNESS, Y: BODY_SIZE_Y - WALL_THICKNESS}, BODY_CURVE)

	walls = sdf.Difference2D(walls, body)

	// left and right side
	lrCutout := trapezoid(v2.Vec{X: (BODY_SIZE_Y - 20) / 2, Y: WALL_THICKNESS / 4}, -30)
	ltCutout := sdf.Transform2D(lrCutout, sdf.Rotate2d(sdf.DtoR(270)))
	ltCutout = sdf.Transform2D(ltCutout, sdf.Translate2d(v2.Vec{X: -((BODY_SIZE_X / 2) - ltCutout.BoundingBox().Max.X), Y: BODY_SIZE_Y / 4}))
	lbCutout := sdf.Transform2D(ltCutout, sdf.MirrorX())
	rtCutout := sdf.Transform2D(ltCutout, sdf.MirrorY())
	rbCutout := sdf.Transform2D(lbCutout, sdf.MirrorY())
	walls = sdf.Difference2D(walls, ltCutout)
	walls = sdf.Difference2D(walls, lbCutout)
	walls = sdf.Difference2D(walls, rtCutout)
	walls = sdf.Difference2D(walls, rbCutout)

	// top and bottom sides
	// TODO: This is producing holes in the render :(
	tbCutout := trapezoid(v2.Vec{X: (BODY_SIZE_X - 20) / 2, Y: WALL_THICKNESS / 4}, -40)
	blCutout := sdf.Transform2D(tbCutout, sdf.Translate2d(v2.Vec{X: -(BODY_SIZE_X / 4), Y: -((BODY_SIZE_Y / 2) - tbCutout.BoundingBox().Max.Y)}))
	brCutout := sdf.Transform2D(blCutout, sdf.MirrorY())
	tlCutout := sdf.Transform2D(blCutout, sdf.MirrorX())
	trCutout := sdf.Transform2D(tlCutout, sdf.MirrorY())
	walls = sdf.Difference2D(walls, blCutout)
	walls = sdf.Difference2D(walls, brCutout)
	walls = sdf.Difference2D(walls, tlCutout)
	walls = sdf.Difference2D(walls, trCutout)

	return walls
}
