package test

import (
	"fmt"
	. "github.com/rinosukmandityo/design-pattern/structural/adapter/model"
	"testing"
)

func TestAdapter(t *testing.T) {
	hole := NewRoundHole(5)
	rpeg := NewConcreteRoundPeg(5)
	fmt.Println(hole.Fits(rpeg)) // true

	small_sqpeg := NewSquarePeg(5)
	large_sqpeg := NewSquarePeg(10)
	// hole.Fits(small_sqpeg) // this won't compile (incompatible types)

	small_sqpeg_adapter := NewSquarePegAdapter(small_sqpeg)
	large_sqpeg_adapter := NewSquarePegAdapter(large_sqpeg)
	fmt.Println(hole.Fits(small_sqpeg_adapter)) // true
	fmt.Println(hole.Fits(large_sqpeg_adapter)) // false
}
