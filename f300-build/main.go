package main

import (
	"log"

	"github.com/deadsy/sdfx/render"
	"github.com/deadsy/sdfx/render/dc"
	"github.com/deadsy/sdfx/sdf"
	v2 "github.com/deadsy/sdfx/vec/v2"
	v3 "github.com/deadsy/sdfx/vec/v3"
)

func main() {
	render.ToSTL(switchBoardMount(), 300, "switchBoardMount.stl", dc.NewDualContouringDefault())
	render.ToSTL(usbCMount(), 300, "usbCMount.stl", dc.NewDualContouringDefault())
	render.ToSTL(usbCCover(), 300, "usbCCover.stl", dc.NewDualContouringDefault())
	render.ToSTL(screenCover(), 300, "screenCover.stl", dc.NewDualContouringDefault())
	render.ToSTL(picoAdapter(), 300, "picoAdapter.stl", dc.NewDualContouringDefault())
}

func picoAdapter() sdf.SDF3 {
	// body
	base2D := sdf.Box2D(v2.Vec{X: 78.9, Y: 39.7}, 0)
	hole2D, _ := sdf.Circle2D(4.9 / 2)

	holes2D := make([]sdf.SDF2, 4)
	holesXOffset, holesYOffset := 4.7, 3.8
	holes2D[0] = sdf.Transform2D(hole2D, sdf.Translate2d(v2.Vec{X: base2D.BoundingBox().Max.X - holesXOffset, Y: base2D.BoundingBox().Max.Y - holesYOffset}))
	holes2D[1] = sdf.Transform2D(hole2D, sdf.Translate2d(v2.Vec{X: -base2D.BoundingBox().Max.X - (-holesXOffset), Y: base2D.BoundingBox().Max.Y - holesYOffset}))
	holes2D[2] = sdf.Transform2D(hole2D, sdf.Translate2d(v2.Vec{X: -base2D.BoundingBox().Max.X - (-holesXOffset), Y: -base2D.BoundingBox().Max.Y - (-holesYOffset)}))
	holes2D[3] = sdf.Transform2D(hole2D, sdf.Translate2d(v2.Vec{X: base2D.BoundingBox().Max.X - holesXOffset, Y: -base2D.BoundingBox().Max.Y - (-holesYOffset)}))
	screwHoles2D := sdf.Union2D(holes2D...)

	base2D = sdf.Difference2D(base2D, screwHoles2D)

	// pico pegs
	m2Diameter := 2.0
	innerCircle, _ := sdf.Circle2D(m2Diameter / 2)
	outerCircle, _ := sdf.Circle2D(m2Diameter * 2 / 2)
	peg2D := sdf.Difference2D(outerCircle, innerCircle)

	pegs2D := make([]sdf.SDF2, 6)
	pegsXDistance, pegsYDistance := 47.0, 11.40
	pegs2D[0] = sdf.Transform2D(peg2D, sdf.Translate2d(v2.Vec{X: pegsXDistance / 2, Y: pegsYDistance / 2}))
	pegs2D[1] = sdf.Transform2D(peg2D, sdf.Translate2d(v2.Vec{X: -pegsXDistance / 2, Y: pegsYDistance / 2}))
	pegs2D[2] = sdf.Transform2D(peg2D, sdf.Translate2d(v2.Vec{X: -pegsXDistance / 2, Y: -pegsYDistance / 2}))
	pegs2D[3] = sdf.Transform2D(peg2D, sdf.Translate2d(v2.Vec{X: pegsXDistance / 2, Y: -pegsYDistance / 2}))
	mount2D := sdf.Union2D(pegs2D...)

	// extrude
	baseZ, pegsZ := 1.6, 3.0
	base := sdf.Extrude3D(base2D, baseZ)
	mount := sdf.Extrude3D(mount2D, pegsZ)
	mount = sdf.Transform3D(mount, sdf.Translate3d(v3.Vec{X: 0, Y: 0, Z: baseZ/2 + pegsZ/2}))

	return sdf.Union3D(base, mount)
}

func screenCover() sdf.SDF3 {
	// cover
	z := 2.2
	holeX, holeY := 32.4, 19.2         // -0.2 for clearance
	screenX, screenY := 23.744, 12.864 // screen V/A
	holderOffsetX, holderOffsetY := 1.75, 2.0

	hole := sdf.Box2D(v2.Vec{X: holeX, Y: holeY}, 0)
	screen := sdf.Box2D(v2.Vec{X: screenX, Y: screenY}, 0)

	body2D := sdf.Difference2D(hole, screen)

	body := sdf.Extrude3D(body2D, z)

	lrEdge2D := sdf.Box2D(v2.Vec{X: (holeX - screenX) / 1.5, Y: holeY - 1.4}, 1)
	lrEdge := sdf.Extrude3D(lrEdge2D, z/2)
	lrEdges := make([]sdf.SDF3, 2)
	lrEdges[0] = sdf.Transform3D(lrEdge, sdf.Translate3d(v3.Vec{X: -holeX/2 - (-(holeX - screenX) / 6), Y: 1.4 / 2, Z: z / 2}))
	lrEdges[1] = sdf.Transform3D(lrEdge, sdf.Translate3d(v3.Vec{X: holeX/2 - ((holeX - screenX) / 6), Y: 1.4 / 2, Z: z / 2}))

	tbEdge2D := sdf.Box2D(v2.Vec{X: holeX, Y: (holeY - screenY) / 2}, 0)
	tbEdge := sdf.Extrude3D(tbEdge2D, z/2)
	tbEdges := make([]sdf.SDF3, 2)
	tbEdges[0] = sdf.Transform3D(tbEdge, sdf.Translate3d(v3.Vec{X: 0, Y: -holeY/2 - (-(holeY - screenY) / 4), Z: z / 2}))
	tbEdges[1] = sdf.Transform3D(tbEdge, sdf.Translate3d(v3.Vec{X: 0, Y: holeY/2 - ((holeY - screenY) / 4), Z: z / 2}))

	body = sdf.Union3D(body, lrEdges[0], lrEdges[1], tbEdges[0], tbEdges[1])

	body = sdf.Transform3D(body, sdf.RotateX(sdf.DtoR(180)))

	// screen mounting

	m2HoleDiameter := 2.0
	PCBHoleX, PCBHoleY := 23.5, 23.8

	m2Hole, _ := sdf.Circle2D(m2HoleDiameter / 2)

	m2Holes := sdf.Union2D(
		sdf.Transform2D(m2Hole, sdf.Translate2d(v2.Vec{X: PCBHoleX / 2, Y: PCBHoleY / 2})),
		sdf.Transform2D(m2Hole, sdf.Translate2d(v2.Vec{X: PCBHoleX / 2, Y: -PCBHoleY / 2})),
		sdf.Transform2D(m2Hole, sdf.Translate2d(v2.Vec{X: -PCBHoleX / 2, Y: -PCBHoleY / 2})),
		sdf.Transform2D(m2Hole, sdf.Translate2d(v2.Vec{X: -PCBHoleX / 2, Y: PCBHoleY / 2})),
	)
	// screen mount pegs

	m2PegFrame2D, _ := sdf.Circle2D(m2HoleDiameter)

	m2PegFrames2D := sdf.Union2D(
		sdf.Transform2D(m2PegFrame2D, sdf.Translate2d(v2.Vec{X: PCBHoleX / 2, Y: -PCBHoleY / 2})),
		sdf.Transform2D(m2PegFrame2D, sdf.Translate2d(v2.Vec{X: -PCBHoleX / 2, Y: -PCBHoleY / 2})),
	)
	m2Pegs2D := sdf.Difference2D(m2PegFrames2D, m2Holes)
	m2Pegs := sdf.Extrude3D(m2Pegs2D, z)
	m2Pegs = sdf.Transform3D(m2Pegs, sdf.Translate3d(v3.Vec{X: 0, Y: 1.5, Z: z}))

	body = sdf.Union3D(body, m2Pegs)

	// peg attachers

	pegHolder, err := sdf.Box3D(v3.Vec{X: 2, Y: holeY / 2, Z: z}, 0)
	if err != nil {
		log.Printf("Model error.\n")
	}

	pegHolders := sdf.Union3D(
		sdf.Transform3D(pegHolder, sdf.Translate3d(v3.Vec{X: (holeX / 2) - holderOffsetX, Y: -holeY/2 + holderOffsetY, Z: z})),
		sdf.Transform3D(pegHolder, sdf.Translate3d(v3.Vec{X: (-holeX / 2) - (-holderOffsetX), Y: -holeY/2 + holderOffsetY, Z: z})),
	)

	body = sdf.Union3D(body, pegHolders)
	return body
}

func usbCCover() sdf.SDF3 {

	usbcWidth := 9.15
	usbcHeight := 3.42
	outHoleX, outHoleY := 14.6, 7.2
	catchX, catchY := 17.0, 9.2
	z := 3.6

	usbCcutout2D := sdf.Box2D(v2.Vec{X: usbcWidth, Y: usbcHeight}, 0.25)

	usbAplug2D := sdf.Box2D(v2.Vec{X: outHoleX, Y: outHoleY}, 1)
	usbAplug2D = sdf.Difference2D(usbAplug2D, usbCcutout2D)

	catch2D := sdf.Box2D(v2.Vec{X: catchX, Y: catchY}, 1)
	catch2D = sdf.Difference2D(catch2D, usbCcutout2D)

	usbAplug := sdf.Extrude3D(usbAplug2D, z)
	catch := sdf.Extrude3D(catch2D, z/6)

	catch = sdf.Transform3D(catch, sdf.Translate3d(v3.Vec{X: 0, Y: 0, Z: -z / 2}))

	return sdf.Union3D(usbAplug, catch)
}

func usbCMount() sdf.SDF3 {
	pcbholeSpacing := 0.65 * sdf.MillimetresPerInch
	//pcbholeDiameter := 0.13 * sdf.MillimetresPerInch
	standoffSpacing := 21.4
	standoffToWallFull, standoffToWallHole := 18.0, 13.0
	m25ScrewHoleDiameter := 2.5
	baseZ, pegsZ := 3.4, 2.6

	base2D := sdf.Box2D(v2.Vec{X: standoffToWallFull, Y: standoffSpacing + 6.0}, 0)
	standOffHole2D, _ := sdf.Circle2D(m25ScrewHoleDiameter / 2)
	standOffHoles2D := sdf.Union2D(
		sdf.Transform2D(standOffHole2D, sdf.Translate2d(v2.Vec{X: (standoffToWallFull-standoffToWallHole)/2 - (m25ScrewHoleDiameter / 2), Y: standoffSpacing / 2})),
		sdf.Transform2D(standOffHole2D, sdf.Translate2d(v2.Vec{X: (standoffToWallFull-standoffToWallHole)/2 - (m25ScrewHoleDiameter / 2), Y: -standoffSpacing / 2})),
	)
	base2D = sdf.Difference2D(base2D, standOffHoles2D)

	peg2D, _ := sdf.Circle2D(m25ScrewHoleDiameter)
	pegHoles2D := sdf.Union2D(
		sdf.Transform2D(standOffHole2D, sdf.Translate2d(v2.Vec{X: 0, Y: pcbholeSpacing / 2})),
		sdf.Transform2D(standOffHole2D, sdf.Translate2d(v2.Vec{X: 0, Y: -pcbholeSpacing / 2})),
	)
	pegs2D := sdf.Union2D(
		sdf.Transform2D(peg2D, sdf.Translate2d(v2.Vec{X: 0, Y: pcbholeSpacing / 2})),
		sdf.Transform2D(peg2D, sdf.Translate2d(v2.Vec{X: 0, Y: -pcbholeSpacing / 2})),
	)
	pegs2D = sdf.Difference2D(pegs2D, pegHoles2D)

	pegs := sdf.Extrude3D(pegs2D, baseZ)
	base := sdf.Extrude3D(base2D, pegsZ)

	pegs = sdf.Transform3D(pegs, sdf.Translate3d(v3.Vec{X: -standoffToWallFull/2 - (-m25ScrewHoleDiameter / 4), Y: 0, Z: baseZ/2 + pegsZ/2}))

	return sdf.Union3D(base, pegs)
}

func switchBoardMount() sdf.SDF3 {
	bodyX, bodyY := 57.0, 26.0
	holeDiameter := 2.7
	holeXSpacing, holeYSpacing := 35.3, 18.7
	chocHole := 13.8
	z := 3.0
	LEDMountDiameter, LEDMountWallThickness := 5.9, 1.2
	LEDLegCutoutX, LEDLegCutoutY := 3.1, 1.0

	// base body for part

	body2D := sdf.Box2D(v2.Vec{X: bodyX, Y: bodyY}, 0)

	mount, _ := sdf.Circle2D(holeDiameter / 2)
	mounts := sdf.Union2D(
		sdf.Transform2D(mount, sdf.Translate2d(v2.Vec{X: -holeXSpacing/2 + bodyX/8, Y: -holeYSpacing / 2})),
		sdf.Transform2D(mount, sdf.Translate2d(v2.Vec{X: holeXSpacing/2 + bodyX/8, Y: holeYSpacing / 2})),
	)

	choc := sdf.Box2D(v2.Vec{X: chocHole, Y: chocHole}, 0)
	chocs := sdf.Union2D(
		sdf.Transform2D(choc, sdf.Translate2d(v2.Vec{X: 15, Y: 0})),
		sdf.Transform2D(choc, sdf.Translate2d(v2.Vec{X: -15, Y: 0})),
		choc,
	)
	chocs = sdf.Transform2D(chocs, sdf.Translate2d(v2.Vec{X: -chocHole / 4, Y: 0}))

	LedThroughHole2D := sdf.Box2D(v2.Vec{X: LEDLegCutoutX, Y: LEDLegCutoutY}, 0)
	LedThroughHole2D = sdf.Transform2D(LedThroughHole2D, sdf.Translate2d(v2.Vec{X: bodyX / 2.55, Y: 0}))

	body2D = sdf.Difference2D(body2D, mounts)
	body2D = sdf.Difference2D(body2D, chocs)
	body2D = sdf.Difference2D(body2D, LedThroughHole2D)

	// Pegs to provide spacing for choc switches

	peg2D, _ := sdf.Circle2D(holeDiameter / 1.5)
	pegs2D := sdf.Union2D(
		sdf.Transform2D(peg2D, sdf.Translate2d(v2.Vec{X: -holeXSpacing/2 + bodyX/8, Y: -holeYSpacing / 2})),
		sdf.Transform2D(peg2D, sdf.Translate2d(v2.Vec{X: holeXSpacing/2 + bodyX/8, Y: holeYSpacing / 2})),
	)
	pegs2D = sdf.Difference2D(pegs2D, mounts)

	pegs := sdf.Extrude3D(pegs2D, z/2)
	body := sdf.Extrude3D(body2D, z)

	pegs = sdf.Transform3D(pegs, sdf.Translate3d(v3.Vec{X: 0, Y: 0, Z: z - z/4}))
	body = sdf.Union3D(body, pegs)

	// Mounting for the LED

	LEDInner, _ := sdf.Circle2D(LEDMountDiameter / 2)
	LEDOutter, _ := sdf.Circle2D((LEDMountDiameter + LEDMountWallThickness) / 2)
	LED2D := sdf.Difference2D(LEDOutter, LEDInner)

	LED := sdf.Extrude3D(LED2D, z/2)

	LED = sdf.Transform3D(LED, sdf.Translate3d(v3.Vec{X: bodyX / 2.55, Y: 0, Z: z - z/4}))
	body = sdf.Union3D(body, LED)
	return body
}
