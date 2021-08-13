package api

type Process interface {
	Name() string

	Call(param interface{}) (interface{}, error)
}
