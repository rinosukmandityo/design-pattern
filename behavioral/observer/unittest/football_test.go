package test

import (
	. "github.com/rinosukmandityo/design-pattern/behavioral/observer/model"
	"testing"
)

func TestFootball(t *testing.T) {
	aBall := NewFootball()
	player1 := NewPlayer("player1", aBall)
	player2 := NewPlayer("player2", aBall)
	player3 := NewPlayer("player3", aBall)
	referee1 := NewReferee("referee1", aBall)

	t.Run("AddPlayer", func(t *testing.T) {
		aBall.Attach(player1)
		aBall.Attach(player2)
		aBall.Attach(player3)
	})
	t.Run("AddReferee", func(t *testing.T) {
		aBall.Attach(referee1)
	})

	t.Run("SetPosition1", func(t *testing.T) {
		aBall.SetPosition(NewPosition(1, 2, 3))
	})

	t.Run("RemovePlayer", func(t *testing.T) {
		aBall.Detach(player2)
	})

	t.Run("SetPosition2", func(t *testing.T) {
		aBall.SetPosition(NewPosition(4, 5, 6))
	})

}
