package module

import "github.com/meidomx/msb/api/kern"

func regInst(i interface{}, err error) error {
	if err != nil {
		return err
	}
	RegisterKernInstance(i)
	return nil
}

func InstantiateAggregator(name string, conf map[string]interface{}, o interface{}) error {
	return regInst(factories[kern.AggregatorFactoryType][name].(kern.AggregatorFactory).LoadConfig(conf, o))
}

func InstantiateBinding(name string, conf map[string]interface{}, o interface{}) error {
	return regInst(factories[kern.BindingFactoryType][name].(kern.BindingFactory).LoadConfig(conf, o))
}

func InstantiateRouter(name string, conf map[string]interface{}, o interface{}) error {
	return regInst(factories[kern.RouterFactoryType][name].(kern.RouterFactory).LoadConfig(conf, o))
}

func InstantiateService(name string, conf map[string]interface{}, o interface{}) error {
	return regInst(factories[kern.ServiceFactoryType][name].(kern.ServiceFactory).LoadConfig(conf, o))
}

func InstantiateTransformer(name string, conf map[string]interface{}, o interface{}) error {
	return regInst(factories[kern.TransformerFactoryType][name].(kern.TransformerFactory).LoadConfig(conf, o))
}

func InstantiateSplitter(name string, conf map[string]interface{}, o interface{}) error {
	return regInst(factories[kern.SplitterFactoryType][name].(kern.SplitterFactory).LoadConfig(conf, o))
}
