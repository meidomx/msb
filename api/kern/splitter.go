package kern

const (
	SplitterFactoryType = "@.factory.splitter"
	SplitterType        = "@.instance.splitter"
)

type SplitterFactory interface {
	Name() string

	LoadConfig(map[string]interface{}, interface{}) (Splitter, error)
}

type Splitter interface {
	Name() string

	Split(interface{}) ([]interface{}, error)
}

type DefaultSplitter struct {
	InstName string
}

func (d *DefaultSplitter) Name() string {
	return d.InstName
}

func (d *DefaultSplitter) Split(i interface{}) ([]interface{}, error) {
	panic("implement me")
}

var _ Splitter = &DefaultSplitter{}
