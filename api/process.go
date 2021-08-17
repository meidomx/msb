package api

type Process interface {
	Name() string

	Call(msbCtx MsbContext, param interface{}) (interface{}, error)
}
