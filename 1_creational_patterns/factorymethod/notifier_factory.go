package factorymethod

import (
	"errors"
	"fmt"
)

// Notification represents a notification.
type Notification struct {
	Messsage string
	To       string
}

// All types of notifiers.
const (
	NOTIFIER_SMS = iota
	NOTIFIER_EMAIL
)

// Notifier exposes all methods of a notifier.
type Notifier interface {
	// SendNotification sends a notification.
	SendNotification(*Notification) error
}

type smsNotifier struct {}

// SendNotification sends a notification to a phone number.
func (n *smsNotifier) SendNotification(notification *Notification) error {
	fmt.Printf("message %s was sent to number %s", notification.Messsage, notification.To)
	return nil
}

type emailNotifier struct {}

// SendNotification sends a notification to a phone number.
func (n *emailNotifier) SendNotification(notification *Notification) error {
	fmt.Printf("message %s was sent to email %s", notification.Messsage, notification.To)
	return nil
}

// Factory Method
func GetNotifier(typ int) (Notifier, error) {
	switch typ {
	case NOTIFIER_SMS: 
		return new(smsNotifier), nil
	case NOTIFIER_EMAIL:
		return new(emailNotifier), nil
	default:
		return nil, errors.New(fmt.Sprintf("notifier type %d has not been supported yet.", typ))
	}
}