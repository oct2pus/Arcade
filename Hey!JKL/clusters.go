package main

import (
	"github.com/deadsy/sdfx/sdf"
)

const (
	KEY_HOLE_SIZE   = 14.2
	KEY_SPACING     = 4.8
	BUTTON_DIAMETER = 24.0
)

func buttonRow() sdf.SDF2 {
	keys := make([]sdf.SDF2, 4)
	for i := range keys {
		keys[i] = sdf.Box2D(sdf.V2{X: KEY_HOLE_SIZE, Y: KEY_HOLE_SIZE}, 0)
		keys[i] = sdf.Transform2D(keys[i], sdf.Translate2d(sdf.V2{
			X: (keys[0].BoundingBox().Max.X*2 + KEY_SPACING) * float64(i),
			Y: 0}))
	}
	buttonRow := sdf.Union2D(keys...)
	buttonRow = sdf.Transform2D(buttonRow, sdf.Translate2d(sdf.V2{
		X: -((keys[0].BoundingBox().Max.X * 2.0) * float64(len(keys))) / 2,
		Y: 0}))
	return buttonRow
}

func hjkl() sdf.SDF2 {
	buttonRow := buttonRow()
	return sdf.Transform2D(buttonRow, sdf.Rotate2d(sdf.DtoR(30)))
}

func buttons() sdf.SDF2 {
	buttons := make([]sdf.SDF2, 4)
	for i := range buttons {
		buttons[i], _ = sdf.Circle2D(BUTTON_DIAMETER / 2)
		buttons[i] = sdf.Transform2D(buttons[i], sdf.Translate2d(sdf.V2{
			X: (buttons[0].BoundingBox().Max.X*2 + KEY_SPACING) * float64(i),
			Y: 0}))
	}
	buttons[len(buttons)-1] = sdf.Transform2D(buttons[len(buttons)-1], sdf.Translate2d(sdf.V2{
		X: 0,
		Y: -buttons[len(buttons)-1].BoundingBox().Max.Y}))
	topRow := sdf.Union2D(buttons...)
	bottomRow := sdf.Transform2D(topRow, sdf.Translate2d(sdf.V2{
		X: 0,
		Y: -(topRow.BoundingBox().Max.Y*2 + KEY_SPACING)}))
	rows := sdf.Union2D(topRow, bottomRow)
	rows = sdf.Transform2D(rows, sdf.Translate2d(sdf.V2{
		X: -(buttons[0].BoundingBox().Max.X*8 - KEY_SPACING*2) / 2,
		Y: topRow.BoundingBox().Max.Y + KEY_SPACING/2}))
	return rows
}

func buttonMounts() sdf.SDF2 {
	buttons := make([]sdf.SDF2, 4)
	mounts := make([]sdf.SDF2, 4)
	for i := range buttons {
		buttons[i], _ = sdf.Circle2D(BUTTON_DIAMETER / 2)
		mounts[i] = sdf.Box2D(sdf.V2{X: KEY_HOLE_SIZE, Y: KEY_HOLE_SIZE}, 0)
		buttons[i] = sdf.Transform2D(buttons[i], sdf.Translate2d(sdf.V2{
			X: (buttons[0].BoundingBox().Max.X*2 + KEY_SPACING) * float64(i),
			Y: 0}))
		mounts[i] = sdf.Transform2D(mounts[i], sdf.Translate2d(sdf.V2{
			X: (buttons[0].BoundingBox().Max.X*2 + KEY_SPACING) * float64(i),
			Y: 0}))
	}
	mounts[len(mounts)-1] = sdf.Transform2D(mounts[len(mounts)-1], sdf.Translate2d(sdf.V2{
		X: 0,
		Y: -buttons[len(buttons)-1].BoundingBox().Max.Y}))
	topRow := sdf.Union2D(buttons...)
	topRowMounts := sdf.Union2D(mounts...)
	bottomRowMounts := sdf.Transform2D(topRowMounts, sdf.Translate2d(sdf.V2{
		X: 0,
		Y: -(topRow.BoundingBox().Max.Y*2 + KEY_SPACING)}))
	rows := sdf.Union2D(topRowMounts, bottomRowMounts)
	rows = sdf.Transform2D(rows, sdf.Translate2d(sdf.V2{
		X: -(buttons[0].BoundingBox().Max.X*8 - KEY_SPACING*2) / 2,
		Y: topRow.BoundingBox().Max.Y + KEY_SPACING/2}))
	return rows
}
