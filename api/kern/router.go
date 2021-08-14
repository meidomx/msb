package kern

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

	Route(result interface{}) string
}

type DefaultRouter struct {
	InstName string
}

func (d *DefaultRouter) Name() string {
	return d.InstName
}

func (d *DefaultRouter) Route(result interface{}) string {
	panic("implement me")
}

var _ Router = &DefaultRouter{}
