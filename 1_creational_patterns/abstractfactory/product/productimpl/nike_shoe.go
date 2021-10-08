package productimpl

import (
	"github/quocdaitrn/go-design-patterns/1_creational_patterns/abstractfactory/product"
)

// AdidasShoe represents an Adidas's shoe.
type NikeShoe struct {
	Logo string
	Size uint
}

// NewAdidasShoe creates an Adidas's shoe.
func NewNikeShoe(size uint) product.Shoe {
	return &NikeShoe{
		Logo: "Nike",
		Size: size,
	}
}

// GetLogo returns logo of the shoe.
func (s *NikeShoe) GetLogo() (string, error) {
	return s.Logo, nil
}

// GetLogo returns size of the shoe.
func (s *NikeShoe) GetSize() (uint, error) {
	return s.Size, nil
}