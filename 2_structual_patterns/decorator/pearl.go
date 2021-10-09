package decorator

// Pearl pearl topping
type Pearl struct {
	Ingredient
}

// Cost calculates the cost of a cup of milk tea.
func (p *Pearl) Cost() (float64, error) {
	base, err := p.Ingredient.Cost()
	return base + 10, err
}