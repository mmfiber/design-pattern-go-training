package bridge

import "fmt"

type PrintableComputer interface {
	Print()
	SetPrinter(Printer)
}

type AbstractPrintableComputer struct {
	printer Printer
}

func (c *AbstractPrintableComputer) SetPrinter(printer Printer) {
	c.printer = printer
}

type Windows struct {
	*AbstractPrintableComputer
}

func (w *Windows) Print() {
	fmt.Println("Print request for windows")
	w.printer.PrintFile()
}

type Mac struct {
	*AbstractPrintableComputer
}

func (m *Mac) Print() {
	fmt.Println("Print request for mac")
	m.printer.PrintFile()
}

type Printer interface {
	PrintFile()
}

type Epson struct{}

func (e *Epson) PrintFile() {
	fmt.Println("Printing by a EPSON Printer")
}

type Hp struct{}

func (h *Hp) PrintFile() {
	fmt.Println("Printing by a HP Printer")
}

type MockPrinter struct{}

func (m *MockPrinter) PrintFile() {}

func main() {
	var printableComputer PrintableComputer
	abstractPrintableComputer := &AbstractPrintableComputer{&MockPrinter{}}

	printableComputer = &Windows{abstractPrintableComputer}
	printableComputer.SetPrinter(&Epson{})
	printableComputer.Print()
	printableComputer.SetPrinter(&Hp{})
	printableComputer.Print()

	printableComputer = &Mac{abstractPrintableComputer}
	printableComputer.SetPrinter(&Epson{})
	printableComputer.Print()
	printableComputer.SetPrinter(&Hp{})
	printableComputer.Print()
}

// codes for main package
type Executer struct{}

func NewExecuter() Executer {
	return Executer{}
}

func (e Executer) Label() string {
	return "bridge pattern"
}

func (e Executer) Do() {
	main()
}
