
package abstractfactory

import (
	"testing"

	"github/quocdaitrn/go-design-patterns/1_creational_patterns/abstractfactory/product/productimpl"
)

func TestAdidasFactory(t *testing.T)  {
	
	adidasFactory, err := BuildSportFactory(FACTORY_ADIDAS)
	if err != nil {
		t.Fatal(err)
	}

	shoe, err := adidasFactory.MakeShoe(40)
	if err != nil {
		t.Fatal(err)
	}

	_, ok := shoe.(*productimpl.AdidasShoe)
	if !ok {
		t.Fatal("Struct assertion failed.")
	}
}

func TestNikeFactory(t *testing.T)  {
	
	nikeFactory, err := BuildSportFactory(FACTORY_NIKE)
	if err != nil {
		t.Fatal(err)
	}

	shoe, err := nikeFactory.MakeShoe(42)
	if err != nil {
		t.Fatal(err)
	}

	_, ok := shoe.(*productimpl.NikeShoe)
	if !ok {
		t.Fatal("Struct assertion failed.")
	}
}