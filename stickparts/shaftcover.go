package main

import "github.com/deadsy/sdfx/sdf"

const (
	SHAFT_GIRTH         = 1.6
	SHAFT_OUTTER_CIRCUM = 11.7
	SHAFT_INNER_CIRCUM  = 9.4
	SHAFT_HEIGHT        = 35.0
)

func shaftCover() (sdf.SDF3, error) {
	sleeve, err := sdf.Circle2D(SHAFT_OUTTER_CIRCUM / 2)
	if err != nil {
		return Error(err)
	}
	hole, err := sdf.Circle2D(SHAFT_INNER_CIRCUM / 2)
	if err != nil {
		return Error(err)
	}
	shaft := sdf.Difference2D(sleeve, hole)

	return sdf.Extrude3D(shaft, SHAFT_HEIGHT), nil
}
