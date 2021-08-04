package api

type MsbHandler interface {
	CallProcess(process string, param interface{}) (interface{}, error)
}
