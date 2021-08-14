package module

import (
	"errors"
	"fmt"

	"github.com/meidomx/msb/api"
	"github.com/meidomx/msb/api/kern"
)

func init() {
	processes = make(map[string]api.Process)

	factories = make(map[string]map[string]interface{})
	instances = make(map[string]map[string]interface{})
}

var hs []api.HttpApiSimpleHandler
var jobs []api.SchedulingJob

var factories map[string]map[string]interface{}
var instances map[string]map[string]interface{}

func RegisterHttpApiHandler(h api.HttpApiSimpleHandler) {
	hs = append(hs, h)
}

func GetHttpApiHandlers() []api.HttpApiSimpleHandler {
	return hs
}

func RegisterSchedulingJob(job api.SchedulingJob) {
	jobs = append(jobs, job)
}

func GetSchedulingJob() []api.SchedulingJob {
	return jobs
}

var processes map[string]api.Process

func RegisterProcess(process api.Process) {
	processes[process.Name()] = process
}

func GetProcess(name string) api.Process {
	return processes[name]
}

func GetBindings() map[string]kern.Binding {
	cp := make(map[string]kern.Binding)
	for k, v := range instances[kern.BindingType] {
		cp[k] = v.(kern.Binding)
	}
	return cp
}

func GetBinding(name string) kern.Binding {
	return instances[kern.BindingType][name].(kern.Binding)
}

func GetServices() map[string]kern.Service {
	cp := make(map[string]kern.Service)
	for k, v := range instances[kern.BindingType] {
		cp[k] = v.(kern.Service)
	}
	return cp
}

func GetService(name string) kern.Service {
	return instances[kern.ServiceType][name].(kern.Service)
}

func RegisterKernInstance(i interface{}) {
	switch item := i.(type) {
	case kern.Aggregator:
		sub, ok := instances[kern.AggregatorType]
		if !ok {
			sub = make(map[string]interface{})
			instances[kern.AggregatorType] = sub
		}
		sub[item.Name()] = item
	case kern.Binding:
		sub, ok := instances[kern.BindingType]
		if !ok {
			sub = make(map[string]interface{})
			instances[kern.BindingType] = sub
		}
		sub[item.Name()] = item
	case kern.Router:
		sub, ok := instances[kern.RouterType]
		if !ok {
			sub = make(map[string]interface{})
			instances[kern.RouterType] = sub
		}
		sub[item.Name()] = item
	case kern.Service:
		sub, ok := instances[kern.ServiceType]
		if !ok {
			sub = make(map[string]interface{})
			instances[kern.ServiceType] = sub
		}
		sub[item.Name()] = item
	case kern.Transformer:
		sub, ok := instances[kern.TransformerType]
		if !ok {
			sub = make(map[string]interface{})
			instances[kern.TransformerType] = sub
		}
		sub[item.Name()] = item
	case kern.Splitter:
		sub, ok := instances[kern.SplitterType]
		if !ok {
			sub = make(map[string]interface{})
			instances[kern.SplitterType] = sub
		}
		sub[item.Name()] = item
	default:
		panic(errors.New("unknown instance type:" + fmt.Sprint(i)))
	}
}

func RegisterKernFactory(f interface{}) {
	switch item := f.(type) {
	case kern.AggregatorFactory:
		sub, ok := factories[kern.AggregatorFactoryType]
		if !ok {
			sub = make(map[string]interface{})
			factories[kern.AggregatorFactoryType] = sub
		}
		sub[item.Name()] = item
	case kern.BindingFactory:
		sub, ok := factories[kern.BindingFactoryType]
		if !ok {
			sub = make(map[string]interface{})
			factories[kern.BindingFactoryType] = sub
		}
		sub[item.Name()] = item
	case kern.RouterFactory:
		sub, ok := factories[kern.RouterFactoryType]
		if !ok {
			sub = make(map[string]interface{})
			factories[kern.RouterFactoryType] = sub
		}
		sub[item.Name()] = item
	case kern.ServiceFactory:
		sub, ok := factories[kern.ServiceFactoryType]
		if !ok {
			sub = make(map[string]interface{})
			factories[kern.ServiceFactoryType] = sub
		}
		sub[item.Name()] = item
	case kern.TransformerFactory:
		sub, ok := factories[kern.TransformerFactoryType]
		if !ok {
			sub = make(map[string]interface{})
			factories[kern.TransformerFactoryType] = sub
		}
		sub[item.Name()] = item
	case kern.SplitterFactory:
		sub, ok := factories[kern.SplitterFactoryType]
		if !ok {
			sub = make(map[string]interface{})
			factories[kern.SplitterFactoryType] = sub
		}
		sub[item.Name()] = item
	default:
		panic(errors.New("unknown factory type:" + fmt.Sprint(f)))
	}
}
