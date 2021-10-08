package factory

import (
	"github/quocdaitrn/go-design-patterns/1_creational_patterns/abstractfactory/product"
	"github/quocdaitrn/go-design-patterns/1_creational_patterns/abstractfactory/product/productimpl"
)

type nikeFactory struct {}

var nikeInstance *nikeFactory

// GetNikeFactory retrieves an instance of Nike's factory.
func GetNikeFactory() *nikeFactory {
	one.Do(func() {
		if nikeInstance == nil {
			nikeInstance = new(nikeFactory)
		}
	})

	return nikeInstance
}

// MakeShoe creates a Nike's shoe.
func (f *nikeFactory) MakeShoe(size uint) (product.Shoe, error) {
	return productimpl.NewNikeShoe(size), nil
}

// MakeShoe creates a Nike's shirt.
func (f *nikeFactory) MakeShirt(color string) (product.Shirt, error) {
	return productimpl.NewNikeShirt(color), nil
}