package model

import (
	"math"
)

type SquarePeg struct {
	width float64
}

func NewSquarePeg(width float64) *SquarePeg {
	return &SquarePeg{width}
}

func (this SquarePeg) GetWidth() float64 {
	return this.width
}

type SquarePegAdapter struct {
	*SquarePeg
}

func NewSquarePegAdapter(peg *SquarePeg) *SquarePegAdapter {
	return &SquarePegAdapter{peg}
}

func (this SquarePegAdapter) GetRadius() float64 {
	return this.GetWidth() * math.Sqrt(2) / 2
}
