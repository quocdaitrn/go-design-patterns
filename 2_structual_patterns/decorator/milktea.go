package decorator

import (
	"errors"
	"fmt"
)

// All sizes of Milktea's cup.
const (
	SIZE_S = iota
	SIZE_M
	SIZE_L
)

// MilkTea concrete of the component.
type MilkTea struct {
	Size        int
	Description string
}

// Cost the cost of Milktea's cup.
func (t *MilkTea) Cost() (float64, error) {
	switch t.Size {
	case SIZE_S:
		return 30, nil
	case SIZE_M:
		return 40, nil
	case SIZE_L:
		return 50, nil
	default:
		return 0, errors.New(fmt.Sprintf("unknown size %d", t.Size))
	}
}
