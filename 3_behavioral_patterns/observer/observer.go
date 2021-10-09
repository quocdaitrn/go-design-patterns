package observer

import (
	"reflect"
	"sync"
)

type Observer interface {
	Notify(string)
}

type Publisher interface {
	Subcribe(o Observer)
	Unsubcribe(o Observer)
	Publish()
}

type publisher struct {
	observers []Observer
	mutext sync.Mutex	
}

func (p *publisher) Subcribe(o Observer) {
	p.mutext.Lock()
	defer p.mutext.Unlock()
	
	p.observers = append(p.observers, o)
}

func (p *publisher) Unsubcribe(o Observer) {
	var idx int
	for i, ob := range p.observers {
		if o == ob {
			idx = i
			break
		}
	}

	p.observers = append(p.observers[:idx], p.observers[idx+1:]...)
}

func (p *publisher) Publish(message string) {
	for _, o := range p.observers {
		o.Notify(message)
	}
}