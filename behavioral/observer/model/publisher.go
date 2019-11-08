package model

type Publisher interface {
	Attach(observer Observer)
	Detach(observer Observer)
	Notify()
}

type DefaultPublisher struct {
	observers []Observer
}

func NewDefaultPublisher() *DefaultPublisher {
	return &DefaultPublisher{[]Observer{}}
}

func (this *DefaultPublisher) Attach(observer Observer) {
	this.observers = append(this.observers, observer)
}

func (this *DefaultPublisher) Detach(observer Observer) {
	for i, obs := range this.observers {
		if obs == observer {
			this.observers = append(this.observers[:i], this.observers[i+1:]...)
		}
	}
}

func (this DefaultPublisher) Notify() {
	for _, obs := range this.observers {
		obs.Update()
	}
}

func (this DefaultPublisher) CountObserver() int {
	return len(this.observers)
}

type ConcretePublisher struct {
	*DefaultPublisher
	message string
}

func NewConcretePublisher() *ConcretePublisher {
	return &ConcretePublisher{NewDefaultPublisher(), ""}
}

func (this ConcretePublisher) GetMessage() string {
	return this.message
}

func (this *ConcretePublisher) SetMessage(m string) {
	this.message = m
	this.Notify()
}
