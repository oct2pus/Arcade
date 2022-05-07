package main

import (
	"github.com/deadsy/sdfx/obj"
	"github.com/deadsy/sdfx/sdf"
)

const BALLTOP_CIRCUM = 35.0

func balltop() (sdf.SDF3, error) {
	cutoutHeight := 6.0

	insert, err := Insert()
	if err != nil {
		return Error(err)
	}
	sphere, err := sdf.Sphere3D(BALLTOP_CIRCUM / 2)
	if err != nil {
		return Error(err)
	}
	cutout, err := sdf.Box3D(sdf.V3{X: BALLTOP_CIRCUM, Y: BALLTOP_CIRCUM, Z: float64(cutoutHeight)}, 0)
	if err != nil {
		return Error(err)
	}
	chamfer, err := obj.ChamferedHole3D(cutoutHeight, 9.8/1.35, cutoutHeight)
	if err != nil {
		return Error(err)
	}
	cutoutShift := (BALLTOP_CIRCUM - cutoutHeight) / 2.0
	cutout = sdf.Transform3D(cutout, sdf.Translate3d(sdf.V3{X: 0, Y: 0, Z: -cutoutShift}))
	sphere = sdf.Difference3D(sphere, cutout)

	chamferShift := (BALLTOP_CIRCUM - cutoutHeight) / 2.0
	chamfer = sdf.Transform3D(chamfer, sdf.Translate3d(sdf.V3{X: 0, Y: 0, Z: -chamferShift}))
	sphere = sdf.Union3D(sphere, chamfer)

	insertShift := (BALLTOP_CIRCUM - INSERT_SIZE_HEIGHT) / 2.0
	insert = sdf.Transform3D(insert, sdf.Translate3d(sdf.V3{X: 0, Y: 0, Z: -insertShift}))
	sphere = sdf.Difference3D(sphere, insert)

	return sphere, nil
}
