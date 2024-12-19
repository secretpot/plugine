package base

type Const struct {
	Name string
	Settings interface{}
}

type Plugin struct {
	Before func()
	After func()
	Const
}