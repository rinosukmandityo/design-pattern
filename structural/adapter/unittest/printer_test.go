package test

import (
	. "github.com/rinosukmandityo/design-pattern/structural/adapter/model"
	"testing"
)

func TestAdapter(t *testing.T) {
	adapter := NewPrinterAdapter(NewConcreteLegacyPrinter(), "Hello World!")
	adapter.PrintStored()
	adapter = NewPrinterAdapter(nil, "Hello World!")
	adapter.PrintStored()
}
