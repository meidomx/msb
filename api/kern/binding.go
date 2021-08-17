package kern

import "github.com/meidomx/msb/api"

const (
	BindingFactoryType = "@.factory.binding"
	BindingType        = "@.instance.binding"
	BindingUsageScope  = "init"
)

type BindingFactory interface {
	Name() string

	LoadConfig(map[string]interface{}, interface{}) (Binding, error)
}

type Binding interface {
	Name() string

	Bind(msbCtx api.MsbContext, parameter interface{}) (interface{}, error)
}

type DefaultBinding struct {
	InstName string
}

func (d *DefaultBinding) Name() string {
	return d.InstName
}

func (d *DefaultBinding) Bind(msbCtx api.MsbContext, parameter interface{}) (interface{}, error) {
	panic("implement me")
}

var _ Binding = &DefaultBinding{}
