package main

import (
	"github.com/deadsy/sdfx/sdf"
	v2 "github.com/deadsy/sdfx/vec/v2"
)

const (
	BODY_SIZE_X = 300
	BODY_SIZE_Y = 220
	BODY_SIZE_Z = 2 + 0 + 0 //Top + Walls + Base
)

func topPlane() sdf.SDF2 {
	top := sdf.Box2D(v2.Vec{X: 300, Y: 220}, 10)
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
