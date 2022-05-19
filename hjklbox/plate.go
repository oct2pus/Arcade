package main

import "github.com/deadsy/sdfx/sdf"

const (
	PLATE_THICKNESS   = 2.825
	TOP_THICKNESS     = 3.175 // 1/8th inch for possible thin acrylic top
	BOTTOM_THICKNESS  = 16 - PLATE_THICKNESS - TOP_THICKNESS
	PLATE_WIDTH       = 218
	PLATE_HEIGHT      = 130
	TOLERANCE         = 8
	CABLE_HEAD_HEIGHT = 6
	CABLE_HEAD_WIDTH  = 10
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

	corners := cornerHoles(plate2D, screwHole())
	plate2D = sdf.Difference2D(plate2D, corners)

	length := lengthHoles(plate2D, smallScrewHole())
	plate2D = sdf.Difference2D(plate2D, length)

	return sdf.Extrude3D(plate2D, PLATE_THICKNESS)
}

func top() sdf.SDF3 {
	top2D := sdf.Box2D(sdf.V2{X: PLATE_WIDTH, Y: PLATE_HEIGHT}, 5)
	hjkl := hjkl()
	hjkl = sdf.Transform2D(hjkl, sdf.Translate2d(sdf.V2{X: top2D.BoundingBox().Max.X / 2, Y: top2D.BoundingBox().Max.Y / 6}))
	top2D = sdf.Difference2D(top2D, hjkl)

	buttons := buttons()
	buttons = sdf.Transform2D(buttons, sdf.Translate2d(sdf.V2{X: -top2D.BoundingBox().Max.X / 2.5, Y: top2D.BoundingBox().Max.Y / 3}))
	top2D = sdf.Difference2D(top2D, buttons)

	buttonRow := buttonRow()
	buttonRow = sdf.Transform2D(buttonRow, sdf.Translate2d(sdf.V2{X: top2D.BoundingBox().Max.X / 2, Y: top2D.BoundingBox().Max.Y / 1.25}))
	top2D = sdf.Difference2D(top2D, buttonRow)

	corners := cornerHoles(top2D, screwHole())
	top2D = sdf.Difference2D(top2D, corners)

	length := lengthHoles(top2D, smallScrewHole())
	top2D = sdf.Difference2D(top2D, length)

	return sdf.Extrude3D(top2D, TOP_THICKNESS)
}

func bottom() sdf.SDF3 {
	cavity2D := sdf.Box2D(sdf.V2{X: PLATE_WIDTH - TOLERANCE, Y: PLATE_HEIGHT - TOLERANCE}, 5)
	walls2D := sdf.Box2D(sdf.V2{X: PLATE_WIDTH, Y: PLATE_HEIGHT}, 5)
	bottom2D := sdf.Difference2D(walls2D, cavity2D)

	cornerScrewHolder, _ := sdf.Circle2D((M4_SCREW_HOLE_DIAMETER + 8) / 2)
	cornerScrewHolders := cornerHoles(bottom2D, cornerScrewHolder)
	cornerScrews := cornerHoles(bottom2D, screwHole())
	bottom2D = sdf.Union2D(bottom2D, cornerScrewHolders)
	bottom2D = sdf.Difference2D(bottom2D, cornerScrews)

	peg, _ := sdf.Circle2D(2.4 / 2) // M2 screw
	pegs := make([]sdf.SDF2, 4)
	for i := range pegs {
		pegs[i] = peg
	}
	pegs[0] = sdf.Transform2D(pegs[0], sdf.Translate2d(sdf.V2{X: 48.26 / 2, Y: 11.4 / 2}))   //pico mounting hole spacing
	pegs[1] = sdf.Transform2D(pegs[1], sdf.Translate2d(sdf.V2{X: 48.26 / 2, Y: -11.4 / 2}))  //pico mounting hole spacing
	pegs[2] = sdf.Transform2D(pegs[2], sdf.Translate2d(sdf.V2{X: -48.26 / 2, Y: -11.4 / 2})) //pico mounting hole spacing
	pegs[3] = sdf.Transform2D(pegs[3], sdf.Translate2d(sdf.V2{X: -48.26 / 2, Y: 11.4 / 2}))  //pico mounting hole spacing
	mount := sdf.Union2D(pegs...)
	mount = sdf.Transform2D(mount, sdf.Translate2d(sdf.V2{X: bottom2D.BoundingBox().Max.X / 2, Y: -bottom2D.BoundingBox().Max.X / 3}))
	bottom2D = sdf.Union2D(bottom2D, mount)
	bottom := sdf.Extrude3D(bottom2D, BOTTOM_THICKNESS)
	floor := sdf.Extrude3D(cavity2D, PLATE_THICKNESS)
	floor = sdf.Transform3D(floor, sdf.Translate3d(sdf.V3{X: 0, Y: 0, Z: -(bottom.BoundingBox().Max.Z/2 + floor.BoundingBox().Max.Z/2)}))
	return sdf.Union3D(bottom, floor)
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
