package bridge

// Computer the Abstraction.
type Computer interface {
	Print() string
}

// Printer the Implementor.
type Printer interface {
	PrintFile() string
}

// Mac a RefinedAbstraction of Computer.
type Mac struct {
	Printer
}

func (m *Mac) Print() string {
	return m.Printer.PrintFile()
}

// Window a RefinedAbstraction of Computer.
type Window struct {
	Printer
}

func (w *Window) Print() string {
	return w.Printer.PrintFile()
}

// EPSON an implementation of Implementor
type EPSON struct {}

func (e *EPSON) PrintFile() string {
	return "File is printed from EPSON Printer"
}

// HP an implementation of Implementor
type HP struct {}

func (h *HP) PrintFile() string {
	return "File is printed from HP Printer"
}

