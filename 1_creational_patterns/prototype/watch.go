package prototype

// Watch exposes all method of a watch.
type Watch interface {
	GetModel() (string, error)
	SetModel(string) error
	Clone() (Watch, error)
}
