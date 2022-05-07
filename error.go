package main

import (
	"github.com/deadsy/sdfx/sdf"
	"github.com/golang/freetype/truetype"
)

func Error(err error) (sdf.SDF3, error) {
	font, _ := truetype.Parse([]byte("./Hack-Regular.ttf"))
	fontSize := 16.0
	height := 2.0
	message := sdf.NewText(err.Error())
	text, _ := sdf.TextSDF2(font, message, fontSize)
	return sdf.Extrude3D(text, height), err
}
