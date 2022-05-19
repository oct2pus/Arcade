package main

import (
	"log"

	"github.com/deadsy/sdfx/sdf"
)

func topPlane() sdf.SDF2 {
	top := sdf.Box2D(sdf.V2{X: 330, Y: 200}, 10)
	joystick := joystick()
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

func splitPlane() map[string]sdf.SDF2 {
	planes := make(map[string]sdf.SDF2)
	original := topPlane()
	x, y := original.BoundingBox().Max.X*2, original.BoundingBox().Max.Y*2
	cutout := sdf.Box2D(sdf.V2{X: x, Y: y}, 0)
	cOLeft := sdf.Transform2D(cutout, sdf.Translate2d(sdf.V2{X: x / 2, Y: 0}))
	cORight := sdf.Transform2D(cutout, sdf.Translate2d(sdf.V2{X: -x / 2, Y: 0}))
	planes["left"] = sdf.Difference2D(original, cOLeft)
	planes["right"] = sdf.Difference2D(original, cORight)
	return planes
}

const M4_SCREW_DIAMETER = 4
const JOYSTICK_TOP_PANEL_HOLE_DIAMETER = 24

// https://support.focusattack.com/hc/en-us/articles/360015744451-Sanwa-JLF-P1-Mounting-Plate-Measurements
// reference for screw hole mounting points

func joystick() sdf.SDF2 {
	holes := make([]sdf.SDF2, 4)
	for i := range holes {
		holes[i], _ = sdf.Circle2D(M4_SCREW_DIAMETER / 2)
	}
	holes[0] = sdf.Transform2D(holes[0], sdf.Translate2d(sdf.V2{X: 36.5, Y: 20}))
	holes[1] = sdf.Transform2D(holes[1], sdf.Translate2d(sdf.V2{X: -36.5, Y: 20}))
	holes[2] = sdf.Transform2D(holes[2], sdf.Translate2d(sdf.V2{X: 36.5, Y: -20}))
	holes[3] = sdf.Transform2D(holes[3], sdf.Translate2d(sdf.V2{X: -36.5, Y: -20}))
	joystickHole, _ := sdf.Circle2D(JOYSTICK_TOP_PANEL_HOLE_DIAMETER / 2)
	holes = append(holes, joystickHole)
	return sdf.Union2D(holes...)
}

const BUTTON_ROW_DIAMETER = 30

// referenced from http://www.slagcoin.com/joystick/layout.html

func buttonRows() sdf.SDF2 {
	buttons := make([]sdf.SDF2, 4)
	for i := range buttons {
		buttons[i], _ = sdf.Circle2D(BUTTON_ROW_DIAMETER / 2)
	}
	buttons[1] = sdf.Transform2D(buttons[1], sdf.Translate2d(sdf.V2{X: 31.25, Y: 9 + 9}))
	buttons[2] = sdf.Transform2D(buttons[2], sdf.Translate2d(sdf.V2{X: 31.25 + 35, Y: 9}))
	buttons[3] = sdf.Transform2D(buttons[3], sdf.Translate2d(sdf.V2{X: 31.25 + 35 + 35, Y: 0}))
	bottomRow := sdf.Union2D(buttons...)
	topRow := bottomRow
	topRow = sdf.Transform2D(topRow, sdf.Translate2d(sdf.V2{X: 0, Y: 9 + 9 + 21}))

	buttonRows := sdf.Union2D(topRow, bottomRow)
	// this is done to recenter these parts for easier work on the topPlane
	buttonRows = sdf.Transform2D(buttonRows, sdf.Translate2d(sdf.V2{X: -(31.25 + 35 + 35) / 2, Y: -(9 + 9 + 21 + 9 + 9) / 2}))
	return buttonRows
}

const FUNCTION_ROW_DIAMETER = 24

func functionRow() sdf.SDF2 {
	spacing := 30.0
	buttonCount := 4
	buttons := make([]sdf.SDF2, buttonCount)
	for i := range buttons {
		buttons[i], _ = sdf.Circle2D(FUNCTION_ROW_DIAMETER / 2)
		buttons[i] = sdf.Transform2D(buttons[i], sdf.Translate2d(sdf.V2{X: spacing * float64(i), Y: 0}))
	}
	functionRow := sdf.Union2D(buttons...)
	functionRow = sdf.Transform2D(functionRow, sdf.Translate2d(sdf.V2{X: -(spacing * float64(buttonCount) / 2), Y: 0}))
	return functionRow
}

func loggedMovement(input sdf.SDF2, displacement sdf.V2, label string) sdf.SDF2 {
	output := sdf.Transform2D(input, sdf.Translate2d(displacement))
	log.Printf("Moving %v by X: %v, Y: %v\n", label, displacement.X, displacement.Y)
	return output
}
