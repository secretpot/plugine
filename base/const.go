package base

type Plugin struct {
	Name string
	Settings interface{}

	Before func() error
	After func() error
	Callback func() error
}