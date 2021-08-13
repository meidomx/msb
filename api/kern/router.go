package kern

import "github.com/meidomx/msb/api"

const (
	RouterFactoryType = "@.factory.router"
	RouterType        = "@.instance.router"
)

type RouterFactory interface {
	Name() string

	LoadConfig(map[string]interface{}, interface{}) (Router, error)
}

type Router interface {
	Name() string

	Route(result api.HandleResult) string
}

type DefaultRouter struct {
	InstName string
}

func (d *DefaultRouter) Name() string {
	return d.InstName
}

func (d *DefaultRouter) Route(result api.HandleResult) string {
	panic("implement me")
}

var _ Router = &DefaultRouter{}
