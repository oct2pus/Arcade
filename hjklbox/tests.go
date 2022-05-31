package main

import "github.com/deadsy/sdfx/sdf"

// 0 is perfect
func usbHoleTest() sdf.SDF3 {
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
}

// i measured 45.6 between holes instead of 47mm, lets try some sizes I guess?
func pegholeTest(pegDistances sdf.V2) sdf.SDF3 {
	plane2D := sdf.Box2D(sdf.V2{X: 50, Y: 30}, 0)
	pegHole, _ := sdf.Circle2D(M2_SCREW_HOLE_DIAMETER / 2)
	pegHoles := make([]sdf.SDF2, 4)
	for i := range pegHoles {
		pegHoles[i] = pegHole
	}
	pegHoles[0] = sdf.Transform2D(pegHoles[0], sdf.Translate2d(sdf.V2{X: pegDistances.X / 2, Y: pegDistances.Y / 2}))   //pico mounting hole spacing
	pegHoles[1] = sdf.Transform2D(pegHoles[1], sdf.Translate2d(sdf.V2{X: pegDistances.X / 2, Y: -pegDistances.Y / 2}))  //pico mounting hole spacing
	pegHoles[2] = sdf.Transform2D(pegHoles[2], sdf.Translate2d(sdf.V2{X: -pegDistances.X / 2, Y: -pegDistances.Y / 2})) //pico mounting hole spacing
	pegHoles[3] = sdf.Transform2D(pegHoles[3], sdf.Translate2d(sdf.V2{X: -pegDistances.X / 2, Y: pegDistances.Y / 2}))  //pico mounting hole spacing

	peg, _ := sdf.Circle2D(M2_SCREW_HOLE_DIAMETER)
	pegs := make([]sdf.SDF2, 4)
	for i := range pegs {
		pegs[i] = peg
	}
	pegs[0] = sdf.Transform2D(pegs[0], sdf.Translate2d(sdf.V2{X: pegDistances.X / 2, Y: pegDistances.Y / 2}))   //pico mounting hole spacing
	pegs[1] = sdf.Transform2D(pegs[1], sdf.Translate2d(sdf.V2{X: pegDistances.X / 2, Y: -pegDistances.Y / 2}))  //pico mounting hole spacing
	pegs[2] = sdf.Transform2D(pegs[2], sdf.Translate2d(sdf.V2{X: -pegDistances.X / 2, Y: -pegDistances.Y / 2})) //pico mounting hole spacing
	pegs[3] = sdf.Transform2D(pegs[3], sdf.Translate2d(sdf.V2{X: -pegDistances.X / 2, Y: pegDistances.Y / 2}))  //pico mounting hole spacing

	mountingHoles2D := sdf.Union2D(pegHoles...)
	mounts2D := sdf.Union2D(pegs...)
	mounts2D = sdf.Difference2D(mounts2D, mountingHoles2D)

	mounts := sdf.Extrude3D(mounts2D, PICO_PEG_HEIGHT)
	plane := sdf.Extrude3D(plane2D, PICO_PEG_HEIGHT/2)

	mounts = sdf.Transform3D(mounts, sdf.Translate3d(sdf.V3{X: 0, Y: 0, Z: PICO_PEG_HEIGHT / 4}))

	return sdf.Union3D(mounts, plane)
}

// usb hole is too low, needs to be a bit higher
func usbmountHeightTest() sdf.SDF3 {
	box, _ := sdf.Box3D(sdf.V3{X: 2, Y: 15, Z: BOTTOM_THICKNESS}, 0)
	floor, _ := sdf.Box3D(sdf.V3{X: 12, Y: 15, Z: 1}, 0)

	floor = sdf.Transform3D(floor, sdf.Translate3d(sdf.V3{X: 5, Y: 0, Z: (-BOTTOM_THICKNESS / 2) - (-floor.BoundingBox().Max.Z)}))

	usbPortHole, _ := sdf.Box3D(sdf.V3{X: USB_CONNECTOR_LENGTH + 0.5, Y: USB_CONNECTOR_HEIGHT + 0.5, Z: USB_CONNECTOR_THICKNESS + 0.5}, 0)
	usbPortHole = sdf.Transform3D(usbPortHole, sdf.Translate3d(sdf.V3{X: 0, Y: 0, Z: 0})) // Z is thickness of board
	usbPortHole = sdf.Transform3D(usbPortHole, sdf.RotateZ(sdf.DtoR(90)))

	box = sdf.Difference3D(box, usbPortHole)
	box = sdf.Union3D(box, floor)

	return box
}
