package test

import (
	. "github.com/rinosukmandityo/design-pattern/structural/adapter/model"
	"testing"
)

func TestAdapter(t *testing.T) {
	var animal Animal

	animal = NewCat()
	animal.Move()
	animal = NewCrocodileAdapter()
	animal.Move()
	animal = NewCrocodileAdapter2()
	animal.Move()
}
