package builtin

import (
	"net/http"

	"github.com/meidomx/msb/api"
	"github.com/meidomx/msb/module"
	"github.com/meidomx/msb/utils"

	"github.com/julienschmidt/httprouter"
)

type Modules struct {
}

func (m Modules) Name() string {
	return "builtin.modules"
}

func (m Modules) HttpMethods() []api.HttpMethod {
	return []api.HttpMethod{
		api.HttpMethodGet,
	}
}

func (m Modules) UrlMapping() string {
	return "/builtin/modules"
}

func (m Modules) Handler() httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		writer.WriteHeader(http.StatusOK)

		m := make(map[string]interface{})
		var httpApis []string
		for _, v := range module.GetHttpApiHandlers() {
			httpApis = append(httpApis, v.Name())
		}
		m["http_api"] = httpApis

		_, err := writer.Write(utils.ObjToJson(m))
		if err != nil {
			LOGGER.Error("builtin modules responding fails.", err)
		}
	}
}

var mgr api.HttpApiHandler = Modules{}

func init() {
	module.RegisterHttpApiHandler(mgr)
}
