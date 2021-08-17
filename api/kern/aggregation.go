package kern

import "github.com/meidomx/msb/api"

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

	Aggregate(msbCtx api.MsbContext, inputs ...interface{}) (interface{}, error)
}

type DefaultAggregator struct {
	InstName string
}

func (d *DefaultAggregator) Name() string {
	return d.InstName
}

func (d *DefaultAggregator) Aggregate(msbCtx api.MsbContext, inputs ...interface{}) (interface{}, error) {
	panic("implement me")
}

var _ Aggregator = &DefaultAggregator{}
