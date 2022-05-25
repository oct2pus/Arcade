package main

import "github.com/deadsy/sdfx/sdf"

const (
	PLATE_THICKNESS             = 2.825
	TOP_THICKNESS               = 3.175 // 1/8th inch for possible thin acrylic top
	BOTTOM_THICKNESS            = 16 - PLATE_THICKNESS - TOP_THICKNESS
	PLATE_WIDTH                 = 218
	PLATE_HEIGHT                = 130
	TOLERANCE                   = 8
	CABLE_HEAD_HEIGHT           = 6
	CABLE_HEAD_WIDTH            = 10
	USB_DAUGHTERBOARD_HEIGHT    = 12.6
	USB_DAUGHTERBOARD_LENGTH    = 21.4
	USB_DAUGHTERBOARD_THICKNESS = 4.7 // includes usb jack
	USB_CONNECTOR_LENGTH        = 8.8 // 1.1mm sticks out from daughterboard
	USB_CONNECTOR_HEIGHT        = 7.2
	USB_CONNECTOR_THICKNESS     = 3.2
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

	//	hjkl := buttonRow

	corners := cornerHoles(top2D, M4screwHole())
	top2D = sdf.Difference2D(top2D, corners)

	length := lengthHoles(top2D, M3ScrewHole())
	top2D = sdf.Difference2D(top2D, length)

	return sdf.Extrude3D(top2D, TOP_THICKNESS)
}

func bottom() sdf.SDF3 {
	cavity2D := sdf.Box2D(sdf.V2{X: PLATE_WIDTH - TOLERANCE, Y: PLATE_HEIGHT - TOLERANCE}, 5)
	walls2D := sdf.Box2D(sdf.V2{X: PLATE_WIDTH, Y: PLATE_HEIGHT}, 5)
	bottom2D := sdf.Difference2D(walls2D, cavity2D)

	cornerScrewHolder, _ := sdf.Circle2D((M4_SCREW_HOLE_DIAMETER + 8) / 2)
	cornerScrewHolders := cornerHoles(bottom2D, cornerScrewHolder)
	cornerScrews := cornerHoles(bottom2D, M4screwHole())
	bottom2D = sdf.Union2D(bottom2D, cornerScrewHolders)
	bottom2D = sdf.Difference2D(bottom2D, cornerScrews)
	cavity2D = sdf.Difference2D(cavity2D, cornerScrewHolders)

	pegHole, _ := sdf.Circle2D(M2_SCREW_HOLE_DIAMETER / 2)
	pegHoles := make([]sdf.SDF2, 4)
	for i := range pegHoles {
		pegHoles[i] = pegHole
	}
	pegHoles[0] = sdf.Transform2D(pegHoles[0], sdf.Translate2d(sdf.V2{X: 47 / 2, Y: 11.4 / 2}))   //pico mounting hole spacing
	pegHoles[1] = sdf.Transform2D(pegHoles[1], sdf.Translate2d(sdf.V2{X: 47 / 2, Y: -11.4 / 2}))  //pico mounting hole spacing
	pegHoles[2] = sdf.Transform2D(pegHoles[2], sdf.Translate2d(sdf.V2{X: -47 / 2, Y: -11.4 / 2})) //pico mounting hole spacing
	pegHoles[3] = sdf.Transform2D(pegHoles[3], sdf.Translate2d(sdf.V2{X: -47 / 2, Y: 11.4 / 2}))  //pico mounting hole spacing

	peg, _ := sdf.Circle2D(M2_SCREW_HOLE_DIAMETER)
	pegs := make([]sdf.SDF2, 4)
	for i := range pegs {
		pegs[i] = peg
	}
	pegs[0] = sdf.Transform2D(pegs[0], sdf.Translate2d(sdf.V2{X: 47 / 2, Y: 11.4 / 2}))   //pico mounting hole spacing
	pegs[1] = sdf.Transform2D(pegs[1], sdf.Translate2d(sdf.V2{X: 47 / 2, Y: -11.4 / 2}))  //pico mounting hole spacing
	pegs[2] = sdf.Transform2D(pegs[2], sdf.Translate2d(sdf.V2{X: -47 / 2, Y: -11.4 / 2})) //pico mounting hole spacing
	pegs[3] = sdf.Transform2D(pegs[3], sdf.Translate2d(sdf.V2{X: -47 / 2, Y: 11.4 / 2}))  //pico mounting hole spacing

	mount2D := sdf.Union2D(pegs...)
	mountingHoles2D := sdf.Union2D(pegHoles...)
	mount2D = sdf.Difference2D(mount2D, mountingHoles2D)
	mount2D = sdf.Transform2D(mount2D, sdf.Translate2d(sdf.V2{X: bottom2D.BoundingBox().Max.X / 2, Y: -bottom2D.BoundingBox().Max.X / 3}))
	mount := sdf.Extrude3D(mount2D, 4) // M2x4mm screw holes
	mount = sdf.Transform3D(mount, sdf.Translate3d(sdf.V3{X: 0, Y: 0, Z: -0.1}))
	//	bottom2D = sdf.Union2D(bottom2D, mount)

	bottom := sdf.Extrude3D(bottom2D, BOTTOM_THICKNESS)
	bottom = sdf.Union3D(bottom, mount)

	cavity2D = sdf.Elongate2D(cavity2D, sdf.V2{X: 1, Y: 1})
	floor := sdf.Extrude3D(cavity2D, PLATE_THICKNESS)
	floor = sdf.Transform3D(floor, sdf.Translate3d(sdf.V3{X: 0, Y: 0, Z: -(bottom.BoundingBox().Max.Z*2+floor.BoundingBox().Max.Z*2)/4 - 0.50}))
	bottom = sdf.Union3D(bottom, floor)

	/*cableHole2D := sdf.Box2D(sdf.V2{X: CABLE_HEAD_WIDTH, Y: CABLE_HEAD_HEIGHT}, 0)
	cableHole := sdf.Extrude3D(cableHole2D, TOLERANCE)
	cableHole = sdf.Transform3D(cableHole, sdf.RotateX(sdf.DtoR(90)))
	cableHole = sdf.Transform3D(cableHole, sdf.Translate3d(sdf.V3{X: 0, Y: (bottom.BoundingBox().Max.Y + cableHole.BoundingBox().Max.Y) - cableHole.BoundingBox().Max.Y*2, Z: cableHole.BoundingBox().Max.Z / 3}))
	bottom = sdf.Difference3D(bottom, cableHole)*/

	usbCutout, _ := sdf.Box3D(sdf.V3{X: USB_DAUGHTERBOARD_LENGTH + 0.4, Y: USB_DAUGHTERBOARD_HEIGHT + 0.4, Z: USB_DAUGHTERBOARD_THICKNESS + 3.8}, 0)
	usbCutout = sdf.Transform3D(usbCutout, sdf.Translate3d(sdf.V3{X: 0, Y: (floor.BoundingBox().Max.Y + usbCutout.BoundingBox().Max.Y) - usbCutout.BoundingBox().Max.Y*2 + 2.6, Z: 2})) // 5 for m3x5 screw
	bottom = sdf.Difference3D(bottom, usbCutout)

	usbCutoutPeg2D := M3ScrewHole()                                 // m2 hole
	usbCutoutPeg := sdf.Extrude3D(usbCutoutPeg2D, BOTTOM_THICKNESS) // thickness of daughtboard pcb
	usbCutoutPeg = sdf.Transform3D(usbCutoutPeg, sdf.Translate3d(sdf.V3{X: 7.8, Y: (floor.BoundingBox().Max.Y + usbCutoutPeg.BoundingBox().Max.Y + 2) - usbCutoutPeg.BoundingBox().Max.Y*2, Z: -BOTTOM_THICKNESS / 2}))
	usbCutoutPeg2 := sdf.Transform3D(usbCutoutPeg, sdf.Translate3d(sdf.V3{X: -15.6, Y: 0, Z: 0}))

	usbCutoutPeg = sdf.Union3D(usbCutoutPeg, usbCutoutPeg2)
	bottom = sdf.Difference3D(bottom, usbCutoutPeg)

	usbPortHole, _ := sdf.Box3D(sdf.V3{X: USB_CONNECTOR_LENGTH, Y: USB_CONNECTOR_HEIGHT, Z: USB_CONNECTOR_THICKNESS}, 0)
	usbPortHole = sdf.Transform3D(usbPortHole, sdf.Translate3d(sdf.V3{X: 0, Y: PLATE_HEIGHT / 2, Z: 1.5 / 2})) // Z is thickness of board

	bottom = sdf.Difference3D(bottom, usbPortHole)

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

// 0 is perfect
/*func usbHoleTest() sdf.SDF3 {
	base2D := sdf.Box2D(sdf.V2{X: 50, Y: 10}, 0)
	holeDimensions := sdf.V2{X: USB_CONNECTOR_LENGTH, Y: USB_CONNECTOR_THICKNESS}
	connectors := make([]sdf.SDF2, 5)

	for i := range connectors {
		connectors[i] = sdf.Box2D(holeDimensions, float64(i))
		connectors[i] = sdf.Transform2D(connectors[i], sdf.Translate2d(sdf.V2{X: (connectors[0].BoundingBox().Max.X*2 + 0.8) * float64(i)}))
	}
	holes := sdf.Union2D(connectors...)
	holes = sdf.Transform2D(holes, sdf.Translate2d(sdf.V2{X: -19, Y: 0}))
	base2D = sdf.Difference2D(base2D, holes)
	return sdf.Extrude3D(base2D, 1.2)
} */

// i measured 45.6 between holes instead of 47mm, lets try some sizes I guess?
func pegholeTest() sdf.SDF3 {

	return nil
}

usb hole is too low, needs to be a bit higher
func usbmountheighttest() sdf.SDF3 {

	return nil
}
