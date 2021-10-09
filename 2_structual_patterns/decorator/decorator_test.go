package decorator

import "testing"

func TestDecorator(t *testing.T) {
	milkTea := &MilkTea{
		Size: SIZE_M,
		Description: "A cup with size M of Milk Tea",
	}

	if cost, _ := milkTea.Cost(); cost != 40 {
		t.Errorf("the cost of medium cup of milk tea without topping should be 40, but found %f", cost)
	}

	puddingPearlMilkTea := &Pudding{&Pearl{milkTea}}

	if cost, _ := puddingPearlMilkTea.Cost(); cost != 65 {
		t.Errorf("the cost of medium cup of milk tea with pearl and pudding topping should be 65, but found %f", cost)
	}
}