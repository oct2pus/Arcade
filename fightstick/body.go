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

	screws := screwHoles()
	top = sdf.Difference2D(top, screws)

	return top
}

// wallsPlane produces a 2D top-down image of the fightstick's walls.
// TODO: use trapezoid to create simpler internal corners for screw holes.
func wallsPlane() sdf.SDF2 {
	walls := sdf.Box2D(v2.Vec{X: BODY_SIZE_X, Y: BODY_SIZE_Y}, BODY_CURVE)
	body := sdf.Box2D(v2.Vec{X: BODY_SIZE_X - WALL_THICKNESS, Y: BODY_SIZE_Y - WALL_THICKNESS}, BODY_CURVE)
	walls = sdf.Difference2D(walls, body)

	screwProtrosion := sdf.Box2D(v2.Vec{X: WALL_THICKNESS / 3, Y: WALL_THICKNESS * 2.2}, 0)
	rProt := loggedMovement(screwProtrosion, v2.Vec{X: (BODY_SIZE_X / 2) - (WALL_THICKNESS / 1.6), Y: 0}, "right screw protrosion")
	lProt := loggedMovement(screwProtrosion, v2.Vec{X: -(BODY_SIZE_X / 2) - -(WALL_THICKNESS / 1.6), Y: 0}, "left screw protrosion")
	tProt := sdf.Transform2D(screwProtrosion, sdf.Rotate2d(sdf.DtoR(90)))
	bProt := loggedMovement(tProt, v2.Vec{X: 0, Y: -(BODY_SIZE_Y / 2) - -(WALL_THICKNESS / 1.6)}, "bottom screw protrosion")
	tProt = loggedMovement(tProt, v2.Vec{X: 0, Y: (BODY_SIZE_Y / 2) - (WALL_THICKNESS / 1.6)}, "top screw protrosion")
	walls = sdf.Union2D(walls, rProt, lProt, tProt, bProt)

	screws := screwHoles()
	walls = sdf.Difference2D(walls, screws)

	return walls
}

// screwHoles produces m4 screwHoles along the sides of the piece.
// TODO: Corners need to be adjusted.
func screwHoles() sdf.SDF2 {
	hole, _ := sdf.Circle2D(M4_SCREW_DIAMETER / 2)
	holes := make([]sdf.SDF2, 14) // 1 top + 1 bottom + (1 * 2 corners) + 2 right, + 1 center = 7 for one side, 14 for two sides.
	for i := range holes {
		holes[i] = hole
	}
	centerOffset := 3.0
	sideOffset := 4.0
	buffer := 2.5
	// right side
	holes[0] = sdf.Transform2D(holes[0], sdf.Translate2d(v2.Vec{X: centerOffset, Y: 0}))
	holes[1] = sdf.Transform2D(holes[1], sdf.Translate2d(v2.Vec{X: centerOffset, Y: (BODY_SIZE_Y / 2) - (WALL_THICKNESS / buffer)}))
	holes[2] = sdf.Transform2D(holes[1], sdf.MirrorX())
	holes[3] = sdf.Transform2D(holes[3], sdf.Translate2d(v2.Vec{X: (BODY_SIZE_X / 2) - (WALL_THICKNESS / buffer), Y: (BODY_SIZE_Y / 2) - (WALL_THICKNESS / buffer)}))
	holes[4] = sdf.Transform2D(holes[3], sdf.MirrorX())
	holes[5] = sdf.Transform2D(holes[5], sdf.Translate2d(v2.Vec{X: (BODY_SIZE_X / 2) - (WALL_THICKNESS / buffer), Y: sideOffset}))
	holes[6] = sdf.Transform2D(holes[5], sdf.MirrorX())

	for o := 0; o < len(holes)/2; o++ {
		holes[o+len(holes)/2] = sdf.Transform2D(holes[o], sdf.MirrorY())
	}

	return sdf.Union2D(holes...)
}
