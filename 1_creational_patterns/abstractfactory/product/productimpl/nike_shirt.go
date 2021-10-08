package productimpl

import (
	"github/quocdaitrn/go-design-patterns/1_creational_patterns/abstractfactory/product"
)

// NikeShirt represents an Nike's shirt.
type NikeShirt struct {
	Logo string
	Color string
}

// NewNikeShirt creates an Nike's shirt.
func NewNikeShirt(color string) product.Shirt {
	return &NikeShirt{
		Logo: "Nike",
		Color: color,
	}
}

// GetLogo returns logo of the shirt.
func (s *NikeShirt) GetLogo() (string, error) {
	return s.Logo, nil
}

// GetLogo returns color of the shirt.
func (s *NikeShirt) GetColor() (string, error) {
	return s.Color, nil
}