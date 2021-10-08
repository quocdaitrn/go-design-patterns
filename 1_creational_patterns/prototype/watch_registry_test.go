package prototype

import (
	"testing"
)

func TestGetWatch(t *testing.T) {
	wr := InitWatchRegistry()

	w, err := wr.GetWatch(WATCH_DIGITAL)
	if err != nil {
		t.Fatal(err)
	}

	dw, ok := w.(*DigitalWatch)
	if !ok {
		t.Error("w should be type DigitalWatch")
	}

	if b, _ := dw.IsAlarm(); b != true {
		t.Error("alarm should be true")
	}
}