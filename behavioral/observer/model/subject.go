package model

// Java's abstract classes combine the definition of interface and default implementation
type Subject interface {
	Attach(observer Observer)
	Detach(observer Observer)
	Notify()
}

// default implementation
type DefaultSubject struct {
	observers []Observer
}

func NewDefaultSubject() *DefaultSubject {
	return &DefaultSubject{[]Observer{}}
}

func (this *DefaultSubject) Attach(observer Observer) {
	this.observers = append(this.observers, observer)
}

func (this *DefaultSubject) Detach(observer Observer) {
	for i, obs := range this.observers {
		if obs == observer {
			this.observers = append(this.observers[:i], this.observers[i+1:]...)
			break
		}
	}
}

func (this DefaultSubject) Notify() {
	for _, obs := range this.observers {
		obs.Update()
	}
}

// concrete subject
type Football struct {
	*DefaultSubject
	position *Position
}

func NewFootball() *Football {
	return &Football{NewDefaultSubject(), nil}
}

func (this *Football) GetPosition() *Position {
	return this.position
}

func (this *Football) SetPosition(position *Position) {
	this.position = position
	this.Notify()
}

type Position struct {
	x, y, z int
}

func NewPosition(x, y, z int) *Position {
	return &Position{x, y, z}
}
