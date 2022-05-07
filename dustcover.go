package main

import (
	"github.com/deadsy/sdfx/obj"
	"github.com/deadsy/sdfx/sdf"
)

const DUST_COVER_OUTER_CIRCUM = 37.5
const DUST_COVER_INNER_CIRCUM = 14.0
const DUST_COVER_HEIGHT = 1.48

func dustCover() (sdf.SDF3, error) {
	cutoutHeight := 0.16
	dustCover, err := obj.Pipe3D(DUST_COVER_OUTER_CIRCUM/2, DUST_COVER_INNER_CIRCUM/2, DUST_COVER_HEIGHT+(cutoutHeight*2))
	if err != nil {
		Error(err)
	}
	cutout1, err := sdf.Box3D(sdf.V3{X: DUST_COVER_OUTER_CIRCUM, Y: DUST_COVER_OUTER_CIRCUM, Z: cutoutHeight}, 0)
	if err != nil {
		Error(err)
	}
	cutout2 := cutout1

	cutout1Shift := (DUST_COVER_HEIGHT + cutoutHeight) / 2
	cutout1 = sdf.Transform3D(cutout1, sdf.Translate3d(sdf.V3{X: 0, Y: 0, Z: cutout1Shift}))

	dustCover = sdf.Difference3D(dustCover, cutout1)

	cutout2Shift := DUST_COVER_HEIGHT / 2
	cutout2 = sdf.Transform3D(cutout2, sdf.Translate3d(sdf.V3{X: 0, Y: 0, Z: -cutout2Shift}))

	dustCover = sdf.Difference3D(dustCover, cutout2)

	return dustCover, nil
}
