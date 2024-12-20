package base

type Const struct {
	Name string
	Settings interface{}
}

type Plugin struct {
	Before func() error
	After func() error
	Callback func() error
	Const
}