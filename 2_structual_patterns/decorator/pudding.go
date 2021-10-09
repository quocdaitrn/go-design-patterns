package decorator

// Pudding pudding topping.
type Pudding struct {
	Ingredient
}

// Cost calculates the cost of a cup of milk tea.
func (p *Pudding) Cost() (float64, error) {
	base, err := p.Ingredient.Cost()
	return base + 15, err
}