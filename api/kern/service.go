package kern

const (
	ServiceFactoryType = "@.factory.service"
	ServiceType        = "@.instance.service"
	ServiceUsageScope  = "runtime"
)

type ServiceFactory interface {
	Name() string

	LoadConfig(map[string]interface{}, interface{}) (Service, error)
}

type Service interface {
	Name() string

	Handle(interface{}) (interface{}, error)
}

type DefaultService struct {
	InstName string
}

func (d *DefaultService) Name() string {
	return d.InstName
}

func (d *DefaultService) Handle(i interface{}) (interface{}, error) {
	panic("implement me")
}

var _ Service = &DefaultService{}
