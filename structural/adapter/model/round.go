package model

type RoundHole struct {
	radius float64
}

func NewRoundHole(radius float64) *RoundHole {
	return &RoundHole{radius}
}

func (this RoundHole) GetRadius() float64 {
	return this.radius
}

func (this RoundHole) Fits(peg RoundPeg) bool {
	return this.GetRadius() >= peg.GetRadius()
}

type RoundPeg interface {
	GetRadius() float64
}

type ConcreteRoundPeg struct {
	radius float64
}

func NewConcreteRoundPeg(radius float64) *ConcreteRoundPeg {
	return &ConcreteRoundPeg{radius}
}

func (this ConcreteRoundPeg) GetRadius() float64 {
	return this.radius
}
