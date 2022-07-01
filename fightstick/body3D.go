package main

import (
	"log"

	"github.com/deadsy/sdfx/sdf"
	v2 "github.com/deadsy/sdfx/vec/v2"
	v3 "github.com/deadsy/sdfx/vec/v3"
)

const (
	TOP_HEIGHT    = 3.0
	WALLS_HEIGHT  = 45.0
	BOTTOM_HEIGHT = 3.0
)

// wallFrontRight is the front right wall. This houses the neutrik connector.
// TODO: Test that measurement because this shit will be infuriating if i print it wrong
func wallFrontRight() sdf.SDF3 {
	corner2D := wallCorner()
	corner := sdf.Extrude3D(corner2D, WALLS_HEIGHT)

	neutrik2D, err := sdf.Circle2D(BUTTON24_DIAMETER / 2)
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	m3Screw, err := sdf.Circle2D(3 / 2)
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	neutrik2D = sdf.Union2D(neutrik2D, sdf.Transform2D(m3Screw, sdf.Translate2d(v2.Vec{X: -19 / 2, Y: 24 / 2})))
	neutrik2D = sdf.Union2D(neutrik2D, sdf.Transform2D(m3Screw, sdf.Translate2d(v2.Vec{X: 19 / 2, Y: -24 / 2})))

	neutrik := sdf.Extrude3D(neutrik2D, WALL_THICKNESS)
	neutrik = sdf.Transform3D(neutrik, sdf.RotateY(sdf.DtoR(90)))
	neutrik = sdf.Transform3D(neutrik, sdf.Translate3d(v3.Vec{X: BODY_SIZE_X/3 + (WALL_THICKNESS / 2), Y: BODY_SIZE_Y / 3, Z: 0}))
	corner = sdf.Difference3D(corner, neutrik)

	return corner
}

//wallFrontLeft is the front left wall. This houses 4 24mm buttons.
func wallFrontLeft() sdf.SDF3 {
	corner2D := wallCorner()
	corner2D = sdf.Transform2D(corner2D, sdf.Rotate2d(sdf.DtoR(270)))
	corner := sdf.Extrude3D(corner2D, WALLS_HEIGHT)

	functionButtons := sdf.Extrude3D(functionRow(), WALL_THICKNESS)
	functionButtons = sdf.Transform3D(functionButtons, sdf.RotateX(sdf.DtoR(90)))
	functionButtons = sdf.Transform3D(functionButtons, sdf.Translate3d(v3.Vec{X: BODY_SIZE_X/4.5 + (WALL_THICKNESS / 2), Y: -BODY_SIZE_Y / 2, Z: 0}))
	corner = sdf.Difference3D(corner, functionButtons)

	return corner
}

//wallBackRight is the back right wall.
//TODO: ROTATE
func wallBackRight() sdf.SDF3 {
	corner := wallCorner()
	corner = sdf.Transform2D(corner, sdf.Rotate2d(sdf.DtoR(90)))
	return sdf.Extrude3D(corner, WALLS_HEIGHT)
}

//wallBackLeft is the back left wall.
//TODO: ROTATE
func wallBackLeft() sdf.SDF3 {
	corner := wallCorner()
	corner = sdf.Transform2D(corner, sdf.Rotate2d(sdf.DtoR(180)))
	return sdf.Extrude3D(corner, WALLS_HEIGHT)
}
