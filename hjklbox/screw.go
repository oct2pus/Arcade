package main

import "github.com/deadsy/sdfx/sdf"

const (
	M4_SCREW_HOLE_DIAMETER = 4.5 //m4 screw
	M3_SCREW_HOLE_DIAMETER = 3.4 //m3 screw
	M2_SCREW_HOLE_DIAMETER = 2.4
)

func M4screwHole() sdf.SDF2 {
	hole, _ := sdf.Circle2D(M4_SCREW_HOLE_DIAMETER / 2)
	return hole
}

func M3ScrewHole() sdf.SDF2 {
	hole, _ := sdf.Circle2D(M3_SCREW_HOLE_DIAMETER / 2)
	return hole
}

func M2ScrewHole() sdf.SDF2 {
	hole, _ := sdf.Circle2D(M2_SCREW_HOLE_DIAMETER / 2)
	return hole
}
