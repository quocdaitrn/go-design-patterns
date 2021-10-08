package prototype

import (
	"bytes"
	"encoding/gob"
)

// DigitalWatch represents an digital watch.
type DigitalWatch struct {
	Model string
	Alarm bool
}

func (w *DigitalWatch) GetModel() (string, error) {
	return w.Model, nil
}

func (w *DigitalWatch) SetModel(model string) error {
	w.Model = model
	return nil
}

func (w *DigitalWatch) IsAlarm() (bool, error) {
	return w.Alarm, nil
}

func (w *DigitalWatch) SetAlarm(b bool) error {
	w.Alarm = b
	return nil
}

func (w *DigitalWatch) Clone() (Watch, error) {
	dw := &DigitalWatch{}
	buf := bytes.Buffer{}

	enc := gob.NewEncoder(&buf)
	enc.Encode(w)

	dec := gob.NewDecoder(&buf)
	dec.Decode(dw)

	return dw, nil
}
