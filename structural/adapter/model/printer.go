package model

import (
	"fmt"
)

type LegacyPrinter interface {
	Print(s string)
}

type ConcreteLegacyPrinter struct {
}

func NewConcreteLegacyPrinter() *ConcreteLegacyPrinter {
	return &ConcreteLegacyPrinter{}
}

func (this ConcreteLegacyPrinter) Print(s string) {
	fmt.Println(fmt.Sprintf("Legacy Printer: %s", s))
}

type ModernPrinter interface {
	PrintStored()
}

type PrinterAdapter struct {
	oldPrinter LegacyPrinter
	msg        string
}

func NewPrinterAdapter(old LegacyPrinter, msg string) *PrinterAdapter {
	return &PrinterAdapter{old, msg}
}

func (this PrinterAdapter) PrintStored() {
	if this.oldPrinter != nil {
		this.oldPrinter.Print(fmt.Sprintf("Adapter: %s", this.msg))
	} else {
		fmt.Println(this.msg)
	}
	return
}
