package bridge

import "testing"

func TestBridge(t *testing.T) {
	epson := &EPSON{}
	hp := &HP{}

	mac := &Mac{}

	var aComputer Computer
	aComputer = mac
	mac.Printer = epson

	if result := aComputer.Print(); result != "File is printed from EPSON Printer" {
		t.Errorf("aComputer should print 'File is printed from EPSON Printer', but found %s", result)
	}

	mac.Printer = hp

	if result := aComputer.Print(); result != "File is printed from HP Printer" {
		t.Errorf("aComputer should print 'File is printed from HP Printer', but found %s", result)
	}
}