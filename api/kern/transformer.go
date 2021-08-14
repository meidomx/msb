package kern

const (
	TransformerFactoryType = "@.factory.transformer"
	TransformerType        = "@.instance.transformer"
	TransformerUsageScope  = "runtime"
)

type TransformerFactory interface {
	Name() string

	LoadConfig(map[string]interface{}, interface{}) (Transformer, error)
}

type Transformer interface {
	Name() string

	Transform(input interface{}) (interface{}, error)
}

type DefaultTransformer struct {
	InstName string
}

func (d *DefaultTransformer) Name() string {
	return d.InstName
}

func (d *DefaultTransformer) Transform(input interface{}) (interface{}, error) {
	panic("implement me")
}

var _ Transformer = &DefaultTransformer{}
