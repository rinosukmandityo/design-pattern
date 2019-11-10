package model

import (
	"fmt"
)

type Animal interface {
	Move()
}

type Cat struct {
}

func NewCat() *Cat {
	return &Cat{}
}

func (this *Cat) Move() {
	fmt.Println("a cat is moving")
}

type Crocodile struct {
}

func NewCrocodile() *Crocodile {
	return &Crocodile{}
}

func (this *Crocodile) Slither() {
	fmt.Println("a crocodile is slithering")
}

// embedding
type CrocodileAdapter struct {
	*Crocodile
}

func NewCrocodileAdapter() *CrocodileAdapter {
	return &CrocodileAdapter{new(Crocodile)}
}

func (this *CrocodileAdapter) Move() {
	this.Slither()
}

// composition
type CrocodileAdapter2 struct {
	crocodile *Crocodile
}

func NewCrocodileAdapter2() *CrocodileAdapter2 {
	return &CrocodileAdapter2{new(Crocodile)}
}

func (this *CrocodileAdapter2) Move() {
	this.crocodile.Slither()
}
