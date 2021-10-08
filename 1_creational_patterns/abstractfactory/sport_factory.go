package abstractfactory

import (
	"errors"
	"fmt"

	"github/quocdaitrn/go-design-patterns/1_creational_patterns/abstractfactory/factory"
	"github/quocdaitrn/go-design-patterns/1_creational_patterns/abstractfactory/product"
)

// SportFactory the factory to create sport's products.
type SportFactory interface {
	// MakeShoe creates a shoe.
	MakeShoe(size uint) (product.Shoe, error)

	// MakeShirt creates a shirt.
	MakeShirt(color string) (product.Shirt, error)
}

// Factories are supported.
const (
	FACTORY_ADIDAS = iota
	FACTORY_NIKE
)

// BuildSportFactory build a sport factory by brand.
func BuildSportFactory(brand int) (SportFactory, error) {
	switch brand {
	case FACTORY_ADIDAS:
		return factory.GetAdidasFactory(), nil
	case FACTORY_NIKE:
		return factory.GetNikeFactory(), nil
	default:
		return nil, errors.New(fmt.Sprintf("brand %d have not been supported yet", brand))
	}
}
