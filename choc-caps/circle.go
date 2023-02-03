package main

import (
	"log"

	"github.com/deadsy/sdfx/sdf"
)

type Circle struct {
	size    float64
	rimSize float64
}

// member functions

func (c Circle) Top2D(size float64) sdf.SDF2 {
	circle, err := sdf.Circle2D(size / 2)
	if err != nil {
		log.Fatalln(err)
	}
	return circle
}

func (c Circle) Rim2D(size float64) sdf.SDF2 {
	rimNegative, err := sdf.Circle2D((size - c.rimSize) / 2)
	if err != nil {
		log.Fatalln(err)
	}
	return sdf.Difference2D(c.Top2D(size), rimNegative)
}

func (c Circle) Size() float64 {
	return c.size
}

// new

func circleNew(size, rimSize float64) Circle {
	var c Circle
	c.size = size
	c.rimSize = rimSize
	return c
}
