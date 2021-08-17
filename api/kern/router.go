package kern

import "github.com/meidomx/msb/api"

const (
	RouterFactoryType = "@.factory.router"
	RouterType        = "@.instance.router"
	RouterUsageScope  = "runtime"
)

type RouterFactory interface {
	Name() string

	LoadConfig(map[string]interface{}, interface{}) (Router, error)
}

type Router interface {
	Name() string

	Route(msbCtx api.MsbContext, result interface{}) string
}

type DefaultRouter struct {
	InstName string
}

func (d *DefaultRouter) Name() string {
	return d.InstName
}

func (d *DefaultRouter) Route(msbCtx api.MsbContext, result interface{}) string {
	panic("implement me")
}

var _ Router = &DefaultRouter{}
