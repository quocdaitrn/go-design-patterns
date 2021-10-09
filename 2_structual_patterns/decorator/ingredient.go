package decorator

// Ingredient component.
type Ingredient interface {
	Cost() (float64, error)
}

