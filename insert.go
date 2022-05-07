package main

import (
	"github.com/deadsy/sdfx/sdf"
)

const INSERT_SIZE_CIRCUM = 8.5
const INSERT_SIZE_HEIGHT = 13.0

func Insert() (sdf.SDF3, error) {
	return sdf.Cylinder3D(INSERT_SIZE_HEIGHT, INSERT_SIZE_CIRCUM/2, 0)
}
