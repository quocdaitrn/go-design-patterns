package factorymethod

import "testing"

func TestGetNotifier(t *testing.T) {
	notifier, err := GetNotifier(NOTIFIER_SMS)
	if err != nil {
		t.Fatal("Notifier type 'SMS' must be returned")
	}

	_, ok := notifier.(*smsNotifier)
	if !ok {
		t.Error("smsNotifier type must be returned")
	}
}