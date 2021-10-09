package observer

import "testing"

func TestObserver(t *testing.T) {
	obs := &Subcriber{}
	sub := &Publisher{}

	obs.Register(sub)

	sub.SetState("publisher's state changed")
	if obs.State != "publisher's state changed" {
		t.Errorf("obs's state must be changed")
	}
}