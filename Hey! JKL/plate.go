package main

import "github.com/deadsy/sdfx/sdf"

const (
	PLATE_THICKNESS             = 2.825
	TOP_THICKNESS               = 3.175 // 1/8th inch for possible thin acrylic top
	BOTTOM_THICKNESS            = 20 - PLATE_THICKNESS - TOP_THICKNESS
	PLATE_WIDTH                 = 218.0
	PLATE_HEIGHT                = 130.0
	TOLERANCE                   = 8.0
	CABLE_HEAD_HEIGHT           = 6.0
	CABLE_HEAD_WIDTH            = 10.0
	USB_DAUGHTERBOARD_HEIGHT    = 12.6
	USB_DAUGHTERBOARD_LENGTH    = 21.4
	USB_DAUGHTERBOARD_THICKNESS = 4.7 // includes usb jack
	USB_CONNECTOR_LENGTH        = 8.8 // 1.1mm sticks out from daughterboard
	USB_CONNECTOR_HEIGHT        = 7.2
	USB_CONNECTOR_THICKNESS     = 3.2
	PICO_PEG_HEIGHT             = 4.0
)

func plate() sdf.SDF3 {
	plate2D := sdf.Box2D(sdf.V2{X: PLATE_WIDTH, Y: PLATE_HEIGHT}, 5)

	hjkl := hjkl()
	hjkl = sdf.Transform2D(hjkl, sdf.Translate2d(sdf.V2{X: plate2D.BoundingBox().Max.X / 2, Y: plate2D.BoundingBox().Max.Y / 6}))
	plate2D = sdf.Difference2D(plate2D, hjkl)

	buttonMounts := buttonMounts()
	buttonMounts = sdf.Transform2D(buttonMounts, sdf.Translate2d(sdf.V2{X: -plate2D.BoundingBox().Max.X / 2.5, Y: plate2D.BoundingBox().Max.Y / 3}))
	plate2D = sdf.Difference2D(plate2D, buttonMounts)

	buttonRow := buttonRow()
	buttonRow = sdf.Transform2D(buttonRow, sdf.Translate2d(sdf.V2{X: plate2D.BoundingBox().Max.X / 2, Y: plate2D.BoundingBox().Max.Y / 1.25}))
	plate2D = sdf.Difference2D(plate2D, buttonRow)

	corners := cornerHoles(plate2D, M4screwHole())
	plate2D = sdf.Difference2D(plate2D, corners)

	length := lengthHoles(plate2D, M3ScrewHole())
	plate2D = sdf.Difference2D(plate2D, length)

	return sdf.Extrude3D(plate2D, PLATE_THICKNESS)
}

func top() sdf.SDF3 {
	top2D := sdf.Box2D(sdf.V2{X: PLATE_WIDTH, Y: PLATE_HEIGHT}, 5)

	hjkl := buttonRow()
	hjkl = sdf.Box2D(sdf.V2{X: hjkl.BoundingBox().Max.X*2 + 2, Y: hjkl.BoundingBox().Max.Y*2 + 2}, 1)
	hjkl = sdf.Transform2D(hjkl, sdf.Rotate2d(sdf.DtoR(30)))
	hjkl = sdf.Transform2D(hjkl, sdf.Translate2d(sdf.V2{X: top2D.BoundingBox().Max.X / 2, Y: top2D.BoundingBox().Max.Y / 6}))
	top2D = sdf.Difference2D(top2D, hjkl)

	buttons := buttons()
	buttons = sdf.Transform2D(buttons, sdf.Translate2d(sdf.V2{X: -top2D.BoundingBox().Max.X / 2.5, Y: top2D.BoundingBox().Max.Y / 3}))
	top2D = sdf.Difference2D(top2D, buttons)

	buttonRow := buttonRow()
	buttonRow = sdf.Box2D(sdf.V2{X: buttonRow.BoundingBox().Max.X*2 + 2, Y: buttonRow.BoundingBox().Max.Y*2 + 2}, 1)
	buttonRow = sdf.Transform2D(buttonRow, sdf.Translate2d(sdf.V2{X: top2D.BoundingBox().Max.X / 2, Y: top2D.BoundingBox().Max.Y / 1.25}))
	top2D = sdf.Difference2D(top2D, buttonRow)

	corners := cornerHoles(top2D, M4screwHole())
	top2D = sdf.Difference2D(top2D, corners)

	length := lengthHoles(top2D, M3ScrewHole())
	top2D = sdf.Difference2D(top2D, length)

	return sdf.Extrude3D(top2D, TOP_THICKNESS)
}

func walls() sdf.SDF3 {
	cavity2D := sdf.Box2D(sdf.V2{X: PLATE_WIDTH - TOLERANCE, Y: PLATE_HEIGHT - TOLERANCE}, 5)
	walls2D := sdf.Box2D(sdf.V2{X: PLATE_WIDTH, Y: PLATE_HEIGHT}, 5)
	bottom2D := sdf.Difference2D(walls2D, cavity2D)

	cornerScrewHolder, _ := sdf.Circle2D((M4_SCREW_HOLE_DIAMETER + 8) / 2)
	cornerScrewHolders := cornerHoles(bottom2D, cornerScrewHolder)
	cornerScrews := cornerHoles(bottom2D, M4screwHole())
	bottom2D = sdf.Union2D(bottom2D, cornerScrewHolders)
	bottom2D = sdf.Difference2D(bottom2D, cornerScrews)
	cavity2D = sdf.Difference2D(cavity2D, cornerScrewHolders)
	bottom := sdf.Extrude3D(bottom2D, BOTTOM_THICKNESS-PLATE_THICKNESS)

	usbCutout, _ := sdf.Box3D(sdf.V3{X: USB_DAUGHTERBOARD_LENGTH + 0.4, Y: USB_DAUGHTERBOARD_HEIGHT + 0.4, Z: USB_DAUGHTERBOARD_THICKNESS + 3}, 0)
	usbCutout = sdf.Transform3D(usbCutout, sdf.Translate3d(sdf.V3{X: 0, Y: (cavity2D.BoundingBox().Max.Y + usbCutout.BoundingBox().Max.Y) - usbCutout.BoundingBox().Max.Y*2 + 2.6, Z: -2})) // 5 for m3x5 screw
	bottom = sdf.Difference3D(bottom, usbCutout)

	usbPortHole, _ := sdf.Box3D(sdf.V3{X: USB_CONNECTOR_LENGTH + 0.5, Y: USB_CONNECTOR_HEIGHT + 0.5, Z: USB_CONNECTOR_THICKNESS + 0.5}, 0)
	usbPortHole = sdf.Transform3D(usbPortHole, sdf.Translate3d(sdf.V3{X: 0, Y: PLATE_HEIGHT / 2, Z: -2})) // Z is thickness of board

	bottom = sdf.Difference3D(bottom, usbPortHole)

	return bottom
}

func bottom() sdf.SDF3 {
	bottom2D := sdf.Box2D(sdf.V2{X: PLATE_WIDTH, Y: PLATE_HEIGHT}, 5)
	cornerScrewHolder, _ := sdf.Circle2D((M4_SCREW_HOLE_DIAMETER + 8) / 2)
	cornerScrewHolders := cornerHoles(bottom2D, cornerScrewHolder)
	cornerScrews := cornerHoles(bottom2D, M4screwHole())

	bottom2D = sdf.Union2D(bottom2D, cornerScrewHolders)
	bottom2D = sdf.Difference2D(bottom2D, cornerScrews)

	usbCutoutPeg2D := M3ScrewHole()
	usbCutoutPeg2D = sdf.Transform2D(usbCutoutPeg2D, sdf.Translate2d(sdf.V2{X: 8.3, Y: (bottom2D.BoundingBox().Max.Y + usbCutoutPeg2D.BoundingBox().Max.Y - 3.2) - usbCutoutPeg2D.BoundingBox().Max.Y*2}))
	usbCutoutPeg2D = sdf.Union2D(usbCutoutPeg2D, sdf.Transform2D(usbCutoutPeg2D, sdf.Translate2d(sdf.V2{X: -16.6, Y: 0})))

	bottom2D = sdf.Difference2D(bottom2D, usbCutoutPeg2D)

	bottom := sdf.Extrude3D(bottom2D, PLATE_THICKNESS)
	pegCount := 4
	pegWall, _ := sdf.Circle2D(M2_SCREW_HOLE_DIAMETER)
	pegHole, _ := sdf.Circle2D(M2_SCREW_HOLE_DIAMETER / 2)
	pegHoles := make([]sdf.SDF2, pegCount)
	pegWalls := make([]sdf.SDF2, pegCount)
	pegs := make([]sdf.SDF3, pegCount)
	for i := 0; i < pegCount; i++ {
		pegHoles[i] = pegHole
		pegWalls[i] = pegWall
		pegs[i] = sdf.Extrude3D(sdf.Difference2D(pegWall, pegHole), PICO_PEG_HEIGHT)
	}
	pegs[0] = sdf.Transform3D(pegs[0], sdf.Translate3d(sdf.V3{X: 47.5 / 2, Y: 11.4 / 2, Z: 0})) // approximate spacing of pico holes
	pegs[1] = sdf.Transform3D(pegs[1], sdf.Translate3d(sdf.V3{X: 47.5 / 2, Y: -11.4 / 2, Z: 0}))
	pegs[2] = sdf.Transform3D(pegs[2], sdf.Translate3d(sdf.V3{X: -47.5 / 2, Y: -11.4 / 2, Z: 0}))
	pegs[3] = sdf.Transform3D(pegs[3], sdf.Translate3d(sdf.V3{X: -47.5 / 2, Y: 11.4 / 2, Z: 0}))

	picoMount := sdf.Union3D(pegs...)
	picoMount = sdf.Transform3D(picoMount, sdf.Translate3d(sdf.V3{X: PLATE_WIDTH / 4, Y: -PLATE_HEIGHT / 3, Z: PLATE_THICKNESS}))

	bottom = sdf.Union3D(bottom, picoMount)

	return bottom
}

func cornerHoles(input, hole sdf.SDF2) sdf.SDF2 {
	cornerHoles := make([]sdf.SDF2, 4)
	for i := range cornerHoles {
		cornerHoles[i] = hole
	}
	cornerHoles[0] = sdf.Transform2D(cornerHoles[0], sdf.Translate2d(sdf.V2{
		X: input.BoundingBox().Max.X - TOLERANCE,
		Y: input.BoundingBox().Max.Y - TOLERANCE}))
	cornerHoles[1] = sdf.Transform2D(cornerHoles[1], sdf.Translate2d(sdf.V2{
		X: -(input.BoundingBox().Max.X - TOLERANCE),
		Y: input.BoundingBox().Max.Y - TOLERANCE}))
	cornerHoles[2] = sdf.Transform2D(cornerHoles[2], sdf.Translate2d(sdf.V2{
		X: -(input.BoundingBox().Max.X - TOLERANCE),
		Y: -(input.BoundingBox().Max.Y - TOLERANCE)}))
	cornerHoles[3] = sdf.Transform2D(cornerHoles[3], sdf.Translate2d(sdf.V2{
		X: input.BoundingBox().Max.X - TOLERANCE,
		Y: -(input.BoundingBox().Max.Y - TOLERANCE)}))

	return sdf.Union2D(cornerHoles...)
}

func lengthHoles(input, hole sdf.SDF2) sdf.SDF2 {
	cornerHoles := make([]sdf.SDF2, 2)
	for i := range cornerHoles {
		cornerHoles[i] = hole
	}
	cornerHoles[0] = sdf.Transform2D(cornerHoles[0], sdf.Translate2d(sdf.V2{
		X: 0,
		Y: input.BoundingBox().Max.Y - TOLERANCE}))
	cornerHoles[1] = sdf.Transform2D(cornerHoles[1], sdf.Translate2d(sdf.V2{
		X: 0,
		Y: -(input.BoundingBox().Max.Y - TOLERANCE)}))

	return sdf.Union2D(cornerHoles...)
}
