package main

import (
	"github.com/deadsy/sdfx/sdf"
)

func topPlane() sdf.SDF2 {
	top := sdf.Box2D(sdf.V2{X: 330, Y: 200}, 10)
	joystick := joystick(sdf.V2{X: 84, Y: 10 + 10 + 20}) // as listed on jlfmeasure.jpg
	joystick = sdf.Transform2D(joystick, sdf.Rotate2d(sdf.DtoR(90)))
	joystick = loggedMovement(joystick, sdf.V2{X: -top.BoundingBox().Max.X / 2, Y: top.BoundingBox().Max.Y / 7}, "joystick")
	top = sdf.Difference2D(top, joystick)

	buttons := buttonRows()
	buttons = loggedMovement(buttons, sdf.V2{X: top.BoundingBox().Max.X / 2, Y: top.BoundingBox().Max.Y / 4}, "button cluster")
	top = sdf.Difference2D(top, buttons)

	auxillaryButtons := functionRow()
	// this is a bit ugly
	auxillaryButtons = loggedMovement(auxillaryButtons, sdf.V2{X: -top.BoundingBox().Max.X / 2.4, Y: 4 * (top.BoundingBox().Max.Y / 5)}, "function cluster")
	top = sdf.Difference2D(top, auxillaryButtons)
	return top

}
