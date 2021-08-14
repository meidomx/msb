package kern

const (
	AggregatorFactoryType = "@.factory.aggregator"
	AggregatorType        = "@.instance.aggregator"
	AggregatorUsageScope  = "runtime"
)

type AggregatorFactory interface {
	Name() string

	LoadConfig(map[string]interface{}, interface{}) (Aggregator, error)
}

type Aggregator interface {
	Name() string

	Aggregate(inputs ...interface{}) (interface{}, error)
}

type DefaultAggregator struct {
	InstName string
}

func (d *DefaultAggregator) Name() string {
	return d.InstName
}

func (d *DefaultAggregator) Aggregate(inputs ...interface{}) (interface{}, error) {
	panic("implement me")
}

var _ Aggregator = &DefaultAggregator{}
