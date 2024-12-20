package base

type Attributes struct {
	Name string
	Settings interface{}
}

type Plugin interface {
	Attributes

	Before() error
	After() error 
	Callback() error
}