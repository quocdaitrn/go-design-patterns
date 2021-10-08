package prototype

import (
	"bytes"
	"encoding/gob"
)

// AnalogWatch represents an analog watch.
type AnalogWatch struct {
	Model string
}

func (w *AnalogWatch) GetModel() (string, error) {
	return w.Model, nil
}

func (w *AnalogWatch) SetModel(model string) error {
	w.Model = model
	return nil
}

func (w *AnalogWatch) Clone() (Watch, error) {
	aw := &AnalogWatch{}
	buf := bytes.Buffer{}

	enc := gob.NewEncoder(&buf)
	enc.Encode(w)

	dec := gob.NewDecoder(&buf)
	dec.Decode(aw)
	
	return aw, nil
}