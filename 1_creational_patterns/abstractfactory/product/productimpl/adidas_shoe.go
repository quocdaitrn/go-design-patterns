package productimpl

import (
	"github/quocdaitrn/go-design-patterns/1_creational_patterns/abstractfactory/product"
)

// AdidasShoe represents an Adidas's shoe.
type AdidasShoe struct {
	Logo string
	Size uint
}

// NewAdidasShoe creates an Adidas's shoe.
func NewAdidasShoe(size uint) product.Shoe {
	return &AdidasShoe{
		Logo: "Adidas",
		Size: size,
	}
}

// GetLogo returns logo of the shoe.
func (s *AdidasShoe) GetLogo() (string, error) {
	return s.Logo, nil
}

// GetLogo returns size of the shoe.
func (s *AdidasShoe) GetSize() (uint, error) {
	return s.Size, nil
}