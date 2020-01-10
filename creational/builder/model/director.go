package model

type Director struct {
	b iBuilder
}

func NewDirector(b iBuilder) *Director {
	return &Director{b: b}
}

func (d *Director) SetDirector(b iBuilder) {
	d.b = b
}

func (d *Director) BuildHouse() house {
	d.b.setWindowType()
	d.b.setDoorType()
	d.b.setNumFloor()
	return d.b.getHouse()
}
