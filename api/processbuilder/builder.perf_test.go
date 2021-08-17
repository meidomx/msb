package processbuilder

import (
	"fmt"
	"testing"

	"github.com/meidomx/msb/api"
	"github.com/meidomx/msb/api/kern"
)

type ExampleService struct {
	kern.DefaultService
}

func (e *ExampleService) Name() string {
	return "exampleService"
}

func (e *ExampleService) Handle(msbCtx api.MsbContext, input interface{}) (interface{}, error) {
	return input, nil
}

type ExampleTransformer struct {
	kern.DefaultTransformer
}

func (e *ExampleTransformer) Name() string {
	return "exampleTransformer"
}

func (e *ExampleTransformer) Transform(msbCtx api.MsbContext, input interface{}) (interface{}, error) {
	return input, nil
}

type ExampleSplitter struct {
	kern.DefaultSplitter
}

func (e *ExampleSplitter) Name() string {
	return "exampleSplitter"
}

func (e *ExampleSplitter) Split(msbCtx api.MsbContext, input interface{}) ([]interface{}, error) {
	return []interface{}{input, input, input}, nil
}

type ExampleAggregator struct {
	kern.DefaultAggregator
}

func (e *ExampleAggregator) Name() string {
	return "exampleAggregator"
}

func (e *ExampleAggregator) Aggregate(msbCtx api.MsbContext, inputs ...interface{}) (interface{}, error) {
	return inputs, nil
}

type ExampleRouter struct {
	kern.DefaultRouter
}

func (e *ExampleRouter) Name() string {
	return "exampleRouter"
}

func (e *ExampleRouter) Route(msbCtx api.MsbContext, result interface{}) string {
	return "2"
}

var exampleService kern.Service = new(ExampleService)
var exampleTransformer kern.Transformer = new(ExampleTransformer)
var exampleSplitter kern.Splitter = new(ExampleSplitter)
var exampleAggregator kern.Aggregator = new(ExampleAggregator)
var exampleRouter kern.Router = new(ExampleRouter)

func ExampleProcessBuilder() {

	// simple process
	p := New().
		Stage(New().
			Stage(exampleService.Handle).
			Stage(exampleService.Handle).
			Stage(exampleService.Handle).
			Build()).
		StageN(
			Split(exampleSplitter),
			Functions(
				Function(exampleService),
				Function(exampleService),
				FunctionC(Transform(exampleTransformer), Function(exampleService)),
			),
			Aggregate(exampleAggregator),
		).
		StageB(
			Route(exampleRouter),
			Cases(
				Case("1", Function(exampleService)),
				Case("2", Function(exampleService)),
				Case("3", FunctionC(Transform(exampleTransformer), Function(exampleService))),
			),
		).
		Stage(exampleService.Handle).
		Stage(exampleService.Handle).
		Build()

	r, err := p(nil, "hello")
	fmt.Println(r, err)

	// Output: [hello hello hello] <nil>
}

func BenchmarkProcessBuilder(b *testing.B) {

	// simple process
	p := New().
		Stage(New().
			Stage(exampleService.Handle).
			Stage(exampleService.Handle).
			Stage(exampleService.Handle).
			Build()).
		StageN(
			Split(exampleSplitter),
			Functions(
				Function(exampleService),
				Function(exampleService),
				FunctionC(Transform(exampleTransformer), Function(exampleService)),
			),
			Aggregate(exampleAggregator),
		).
		StageB(
			Route(exampleRouter),
			Cases(
				Case("1", Function(exampleService)),
				Case("2", Function(exampleService)),
				Case("3", FunctionC(Transform(exampleTransformer), Function(exampleService))),
			),
		).
		Stage(exampleService.Handle).
		Stage(exampleService.Handle).
		Build()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := p(nil, "hello")
		if err != nil {
			b.Fatal(err)
		}
	}
}
