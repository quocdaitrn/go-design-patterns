package product

// Shirt exposes all shirt's methods.
type Shirt interface {
	// GetLogo returns logo of the shirt.
	GetLogo() (string, error)

	// GetLogo returns color of the shirt.
	GetColor() (string, error)
}