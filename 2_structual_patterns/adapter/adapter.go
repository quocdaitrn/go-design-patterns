package adapter

import "fmt"

// Computer the target.
type Computer interface {
	InsertInSquarePort()
}

// Mac an implementation of Computer.
type Mac struct {}

func (m *Mac) InsertInSquarePort() {
	fmt.Println("Insert square port into mac machine")
}

// WindowMachineAdapter the adapter.
type WindowMachineAdapter struct {
	WindowMachine	
}

func (m *WindowMachineAdapter) InsertInSquarePort() {
	m.WindowMachine.InsertInCirclePort()
}

// WindowMachine the adaptee.
type WindowMachine struct {}

func (m *WindowMachine) InsertInCirclePort() {
	fmt.Println("Insert circle port into window machine")
}

// Client the client.
type Client struct {}

func (c *Client) InsertSquareUsbInComputer(com Computer) {
	com.InsertInSquarePort()
}