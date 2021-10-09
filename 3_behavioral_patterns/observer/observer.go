package observer

import (
	"fmt"
	"sync"
)

type Observable interface {
	Attach(o Observer)
	Detach(o Observer)
	Notify(obj interface{})
}

type Observer interface {
	Register(Observable)
	Update(Observable, interface{})
}

type Publisher struct {
	State string

	observers []Observer
	mutext sync.Mutex	
}

func (p *Publisher) Attach(o Observer) {
	p.mutext.Lock()
	defer p.mutext.Unlock()
	
	p.observers = append(p.observers, o)
}

func (p *Publisher) Detach(o Observer) {
	var idx int
	for i, ob := range p.observers {
		if o == ob {
			idx = i
			break
		}
	}

	p.observers = append(p.observers[:idx], p.observers[idx+1:]...)
}

func (p *Publisher) Notify(messge interface{}) {
	for _, o := range p.observers {
		o.Update(p, messge)
	}
}

func (p *Publisher) SetState(s string) {
	oldState := p.State
	p.State = s
	if p.State != oldState {
		p.Notify(p.State)
	}
}

type Subcriber struct {
	State string
}

func (s *Subcriber) Register(o Observable) {
	o.Attach(s)
}

func (s *Subcriber) Update(o Observable, message interface{}) {
	s.State = message.(string)
	fmt.Printf("subcriber had received the message: %s", message.(string))
}
