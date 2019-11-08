package test

import (
	. "github.com/rinosukmandityo/design-pattern/behavioral/observer/model"
	"testing"
)

func TestObservers(t *testing.T) {
	pub := NewConcretePublisher()
	obs1 := NewConcreteObserver(1, pub)
	obs2 := NewConcreteObserver(2, pub)
	obs3 := NewConcreteObserver(3, pub)

	t.Run("AddObserver", func(t *testing.T) {
		pub.Attach(obs1)
		pub.Attach(obs2)
		pub.Attach(obs3)
		if pub.CountObserver() != 3 {
			t.Fail()
		}
	})

	t.Run("Notify", func(t *testing.T) {
		pub.SetMessage("Hey There!")
	})

	t.Run("RemoveObserver", func(t *testing.T) {
		pub.Detach(obs2)
	})
	t.Run("NotifyAgain", func(t *testing.T) {
		pub.SetMessage("It's time to say goodbye")
	})
}
