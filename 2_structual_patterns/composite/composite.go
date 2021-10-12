package composite

import (
	"errors"
	"fmt"
)

// Graphic the component.
type Graphic interface {
	Draw()
	GetName() string
}

// Line a leaf.
type Line struct {
	Name string
}

func (l *Line) Draw() {
	fmt.Printf("Drawing line %s\n", l.Name)
}

func (l *Line) GetName() string {
	return l.Name
}

// Rectangle a leaf.
type Rectangle struct {
	Name string
}

func (r *Rectangle) Draw() {
	fmt.Printf("Drawing rectangle %s\n", r.Name)
}

func (r *Rectangle) GetName() string {
	return r.Name
}

// Picture the composite.
type Picture struct {
	Graphics []Graphic
	Name string
}

func (p *Picture) Draw() {
	for _, g := range p.Graphics {
		g.Draw()
	}
}

func (p *Picture) GetName() string {
	return p.Name
}

func (p *Picture) Add(g Graphic) {
	p.Graphics = append(p.Graphics, g)
}

func (p *Picture) Remove(g Graphic) {
	index := -1
	for i, gra := range p.Graphics {
		if gra.GetName() == g.GetName() {
			index = i
			break
		}
	}

	if index < 0 {
		return
	}
	p.Graphics = append(p.Graphics[:index], p.Graphics[index+1:]...)
}

func (p *Picture) GetChild(name string) (Graphic, error) {
	for _, g := range p.Graphics {
		if g.GetName() == name {
			return g, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("can not find the child naming %s", name))
}