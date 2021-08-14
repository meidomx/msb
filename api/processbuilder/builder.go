package processbuilder

import (
	"container/list"
	"errors"

	"github.com/meidomx/msb/api"
	"github.com/meidomx/msb/api/kern"
)

type simpleProcessBuilder struct {
	list *list.List
}

type simpleProcess struct {
	name string
	f    funcOneToOne
}

func (s *simpleProcess) Name() string {
	return s.name
}

func (s *simpleProcess) Call(param interface{}) (interface{}, error) {
	return s.f(param)
}

var _ api.Process = new(simpleProcess)

func New() *simpleProcessBuilder {
	p := new(simpleProcessBuilder)
	p.list = list.New()
	return p
}

func (process *simpleProcessBuilder) Stage(function funcOneToOne) *simpleProcessBuilder {
	process.list.PushBack(function)
	return process
}

func (process *simpleProcessBuilder) StageN(split funcOneToOne, functions []funcOneToOne, aggregate funcOneToOne) *simpleProcessBuilder {
	var stageN funcOneToOne = func(i interface{}) (interface{}, error) {
		sp, err := split(i)
		if err != nil {
			return nil, err
		}
		spList, ok := sp.([]interface{})
		if !ok {
			return nil, errors.New("not a slice after split")
		}
		if len(spList) != len(functions) {
			return nil, errors.New("split function length not match")
		}
		var rlist []interface{}
		for i, v := range spList {
			r, err := functions[i](v)
			if err != nil {
				return nil, err
			}
			rlist = append(rlist, r)
		}
		return aggregate(rlist)
	}
	process.list.PushBack(stageN)
	return process
}

func (process *simpleProcessBuilder) StageB(route funcOneToOne, cases funcOneToOne) *simpleProcessBuilder {
	var stageB funcOneToOne = func(i interface{}) (interface{}, error) {
		r, err := route(i)
		if err != nil {
			return nil, err
		}
		rr, ok := r.(string)
		if !ok {
			return nil, errors.New("router result is not string")
		}
		f, err := cases(rr)
		if err != nil {
			return nil, err
		}
		ff, ok := f.(funcOneToOne)
		if !ok {
			return nil, errors.New("case should be a function")
		}
		return ff(i)
	}
	process.list.PushBack(stageB)
	return process
}

func (process *simpleProcessBuilder) Build() funcOneToOne {
	return func(i interface{}) (interface{}, error) {
		for e := process.list.Front(); e != nil; e = e.Next() {
			f := e.Value.(funcOneToOne)
			r, err := f(i)
			if err != nil {
				return nil, err
			}
			i = r
		}
		return i, nil
	}
}

func (process *simpleProcessBuilder) BuildProcess(name string) api.Process {
	p := new(simpleProcess)
	p.name = name
	p.f = process.Build()
	return p
}

type funcOneToOne func(i interface{}) (interface{}, error)

func Function(service kern.Service) funcOneToOne {
	return func(i interface{}) (interface{}, error) {
		return service.Handle(i)
	}
}

func Transform(transformer kern.Transformer) funcOneToOne {
	return func(i interface{}) (interface{}, error) {
		return transformer.Transform(i)
	}
}

func FunctionC(transform funcOneToOne, function funcOneToOne) funcOneToOne {
	return func(i interface{}) (interface{}, error) {
		r, err := transform(i)
		if err != nil {
			return nil, err
		}
		return function(r)
	}
}

func Functions(funcs ...funcOneToOne) []funcOneToOne {
	return funcs
}

func Split(splitter kern.Splitter) funcOneToOne {
	return func(i interface{}) (interface{}, error) {
		return splitter.Split(i)
	}
}

func Aggregate(aggregator kern.Aggregator) funcOneToOne {
	return func(i interface{}) (interface{}, error) {
		ll := i.([]interface{})
		return aggregator.Aggregate(ll...)
	}
}

func Route(router kern.Router) funcOneToOne {
	return func(i interface{}) (interface{}, error) {
		return router.Route(i), nil
	}
}

type caseStruct struct {
	key      string
	function funcOneToOne
}

func Case(s string, function funcOneToOne) caseStruct {
	return caseStruct{
		key:      s,
		function: function,
	}
}

func Cases(c ...caseStruct) funcOneToOne {
	mapping := make(map[string]funcOneToOne)
	for _, v := range c {
		mapping[v.key] = v.function
	}
	return func(i interface{}) (interface{}, error) {
		f, ok := mapping[i.(string)]
		if !ok {
			return nil, errors.New("cannot find case for:" + i.(string))
		}
		return f, nil
	}
}
