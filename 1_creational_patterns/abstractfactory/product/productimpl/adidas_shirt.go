package productimpl

import (
	"github/quocdaitrn/go-design-patterns/1_creational_patterns/abstractfactory/product"
)

// AdidasShirt represents an Adidas's shirt.
type AdidasShirt struct {
	Logo string
	Color string
}

// NewAdidasShirt creates an Adidas's shirt.
func NewAdidasShirt(color string) product.Shirt {
	return &AdidasShirt{
		Logo: "Adidas",
		Color: color,
	}
}

// GetLogo returns logo of the shirt.
func (s *AdidasShirt) GetLogo() (string, error) {
	return s.Logo, nil
}

// GetColor returns color of the shirt.
func (s *AdidasShirt) GetColor() (string, error) {
	return s.Color, nil
}