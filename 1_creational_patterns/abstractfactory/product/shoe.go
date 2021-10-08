package product

// Shoe exposes all shoe's method.
type Shoe interface {
	// GetLogo returns logo of the shoe.
	GetLogo() (string, error)

	// GetLogo returns size of the shoe.
	GetSize() (uint, error)
}