package model

import (
	"fmt"
)

type Observer interface {
	Update()
}

type Player struct {
	name string
	ball *Football
}

func NewPlayer(name string, ball *Football) *Player {
	return &Player{name, ball}
}
func (this Player) Update() {
	fmt.Println(this.name, "noticed the ball has moved to", this.ball.GetPosition())
}

type Referee struct {
	name string
	ball *Football
}

func NewReferee(name string, ball *Football) *Referee {
	return &Referee{name, ball}
}

func (this *Referee) Update() {
	fmt.Println(this.name, "noticed the ball has moved", this.ball.GetPosition())
}

type ConcreteObserver struct {
	id  int
	pub *ConcretePublisher
}

func NewConcreteObserver(id int, pub *ConcretePublisher) *ConcreteObserver {
	return &ConcreteObserver{id, pub}
}

func (this ConcreteObserver) Update() {
	fmt.Printf("Observer %d: message '%s' received\n", this.id, this.pub.GetMessage())
}
