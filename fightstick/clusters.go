package main

import (
	"log"

	"github.com/deadsy/sdfx/sdf"
	v2 "github.com/deadsy/sdfx/vec/v2"
)

const (
	BUTTON30_DIAMETER      = 30.5
	BUTTON24_DIAMETER      = 24.5
	M3_SCREW_DIAMETER      = 3.2
	M4_SCREW_DIAMETER      = 4
	JOYSTICK_HOLE_DIAMETER = 24
)

// buttonRow is the face button cluster.
// referenced from http://www.slagcoin.com/joystick/layout.html,
// measurements are slant36_l.png, does not consider stick distance.
func buttonRows() sdf.SDF2 {
	buttons := make([]sdf.SDF2, 4)
	for i := range buttons {
		buttons[i], _ = sdf.Circle2D(BUTTON30_DIAMETER / 2)
	}
	buttons[1] = sdf.Transform2D(buttons[1], sdf.Translate2d(v2.Vec{X: 31.25, Y: 9 + 9}))
	buttons[2] = sdf.Transform2D(buttons[2], sdf.Translate2d(v2.Vec{X: 31.25 + 35, Y: 9}))
	buttons[3] = sdf.Transform2D(buttons[3], sdf.Translate2d(v2.Vec{X: 31.25 + 35 + 35, Y: 0}))
	bottomRow := sdf.Union2D(buttons...)
	topRow := bottomRow
	topRow = sdf.Transform2D(topRow, sdf.Translate2d(v2.Vec{X: 0, Y: 9 + 9 + 21}))

	buttonRows := sdf.Union2D(topRow, bottomRow)
	// this is done to recenter these parts for easier work on the topPlane
	//buttonRows = sdf.Transform2D(buttonRows, sdf.Translate2d(v2.Vec{X: -(31.25 + 35 + 35) / 2, Y: -(9 + 9 + 21 + 9 + 9) / 2}))
	buttonRows = sdf.Center2D(buttonRows)

	return buttonRows
}

// functionRow is the start/select/home/capture cluster.
func functionRow() sdf.SDF2 {
	spacing := 30.0
	buttonCount := 4
	buttons := make([]sdf.SDF2, buttonCount)
	for i := range buttons {
		buttons[i], _ = sdf.Circle2D(BUTTON24_DIAMETER / 2)
		buttons[i] = sdf.Transform2D(buttons[i], sdf.Translate2d(v2.Vec{X: spacing * float64(i), Y: 0}))
	}
	functionRow := sdf.Union2D(buttons...)
	//functionRow = sdf.Transform2D(functionRow, sdf.Translate2d(v2.Vec{X: -(spacing * float64(buttonCount) / 2), Y: 0}))
	functionRow = sdf.Center2D(functionRow)
	return functionRow
}

// joystick is a hole for a joystick, with screw mounts.
// https://support.focusattack.com/hc/en-us/articles/360015744451-Sanwa-JLF-P1-Mounting-Plate-Measurements
// reference for screw hole mounting points
func joystick(holeSpacing v2.Vec) sdf.SDF2 {
	holes := make([]sdf.SDF2, 4)
	for i := range holes {
		holes[i], _ = sdf.Circle2D(M4_SCREW_DIAMETER / 2)
	}
	holes[0] = sdf.Transform2D(holes[0], sdf.Translate2d(v2.Vec{X: holeSpacing.X / 2, Y: holeSpacing.Y / 2}))
	holes[1] = sdf.Transform2D(holes[1], sdf.Translate2d(v2.Vec{X: -holeSpacing.X / 2, Y: holeSpacing.Y / 2}))
	holes[2] = sdf.Transform2D(holes[2], sdf.Translate2d(v2.Vec{X: holeSpacing.X / 2, Y: -holeSpacing.Y / 2}))
	holes[3] = sdf.Transform2D(holes[3], sdf.Translate2d(v2.Vec{X: -holeSpacing.X / 2, Y: -holeSpacing.Y / 2}))
	joystickHole, _ := sdf.Circle2D(JOYSTICK_HOLE_DIAMETER / 2)
	holes = append(holes, joystickHole)
	return sdf.Union2D(holes...)
}

// neutrik is a hole for a neutrik connector.
// https://focusattack.com/cliff-electronics-usb-c-to-usb-c-passthrough/
// reference for neutrik connector dimensions
func neutrik() sdf.SDF2 {
	neutrik2D, err := sdf.Circle2D(BUTTON24_DIAMETER / 2)
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	m3Screw, err := sdf.Circle2D(M3_SCREW_DIAMETER / 2)
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	neutrik2D = sdf.Union2D(neutrik2D, sdf.Transform2D(m3Screw, sdf.Translate2d(v2.Vec{X: -19.1 / 2, Y: 24 / 2})))
	neutrik2D = sdf.Union2D(neutrik2D, sdf.Transform2D(m3Screw, sdf.Translate2d(v2.Vec{X: 19.1 / 2, Y: -24 / 2})))
	return neutrik2D
}

// screwHoles produces m4 screwHoles along the sides of the piece.
func screwHoles() sdf.SDF2 {
	hole, _ := sdf.Circle2D(M4_SCREW_DIAMETER / 2)
	holeCenter, _ := sdf.Circle2D(M3_SCREW_DIAMETER / 2)
	holes := make([]sdf.SDF2, 14) // 1 top + 1 bottom + (1 * 2 corners) + 2 right, + 1 center = 7 for one side, 14 for two sides.
	for i := range holes {
		holes[i] = hole
	}
	centerOffset := 3.0
	sideOffset := 4.0
	buffer := 2.5
	// right side
	holes[0] = holeCenter
	holes[0] = sdf.Transform2D(holes[0], sdf.Translate2d(v2.Vec{X: centerOffset, Y: 0}))
	holes[1] = sdf.Transform2D(holes[1], sdf.Translate2d(v2.Vec{X: centerOffset, Y: (BODY_SIZE_Y / 2) - (WALL_THICKNESS / buffer)}))
	holes[2] = sdf.Transform2D(holes[1], sdf.MirrorX())
	holes[3] = sdf.Transform2D(holes[3], sdf.Translate2d(v2.Vec{X: (BODY_SIZE_X / 2) - (WALL_THICKNESS), Y: (BODY_SIZE_Y / 2) - (WALL_THICKNESS)}))
	holes[4] = sdf.Transform2D(holes[3], sdf.MirrorX())
	holes[5] = sdf.Transform2D(holes[5], sdf.Translate2d(v2.Vec{X: (BODY_SIZE_X / 2) - (WALL_THICKNESS / buffer), Y: sideOffset}))
	holes[6] = sdf.Transform2D(holes[5], sdf.MirrorX())

	for o := 0; o < len(holes)/2; o++ {
		holes[o+len(holes)/2] = sdf.Transform2D(holes[o], sdf.MirrorY())
	}

	return sdf.Union2D(holes...)
}
