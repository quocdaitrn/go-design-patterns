package factory

import (
	"github/quocdaitrn/go-design-patterns/1_creational_patterns/abstractfactory/product"
	"github/quocdaitrn/go-design-patterns/1_creational_patterns/abstractfactory/product/productimpl"
)

type adidasFactory struct {}

var adidasInstance *adidasFactory

// GetNikeFactory retrieves an instance of Adidas's factory.
func GetAdidasFactory() *adidasFactory {
	one.Do(func() {
		if adidasInstance == nil {
			adidasInstance = new(adidasFactory)
		}
	})

	return adidasInstance
}

// MakeShoe creates a Adidas's shoe.
func (f *adidasFactory) MakeShoe(size uint) (product.Shoe, error) {
	return productimpl.NewAdidasShoe(size), nil
}

// MakeShoe creates a Adidas's shoe.
func (f *adidasFactory) MakeShirt(color string) (product.Shirt, error) {
	return productimpl.NewAdidasShirt(color), nil
}