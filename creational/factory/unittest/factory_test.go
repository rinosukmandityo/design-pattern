package test

import (
	"fmt"
	"testing"

	. "github.com/rinosukmandityo/design-pattern/creational/factory/model"
)

func TestFactory(t *testing.T) {
	ak47, _ := GetGun("ak47")
	maverick, _ := GetGun("maverick")

	printDetails(ak47)
	printDetails(maverick)
}

func printDetails(g IGun) {
	fmt.Printf("Gun: %s", g.GetName())
	fmt.Println()
	fmt.Printf("Power: %d", g.GetPower())
	fmt.Println()
}
