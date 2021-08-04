package builtin

import (
	"net/http"
	"reflect"

	"github.com/meidomx/msb/api"
	"github.com/meidomx/msb/module"
)

type Modules struct {
}

type ManagementRequest struct {
	Operation string
}

func (this Modules) ContentTypes() (api.RequestFormat, api.ResponseFormat) {
	return api.RequestFormatJson, api.ResponseFormatJson
}

func (this Modules) RequestType() reflect.Type {
	return reflect.TypeOf(ManagementRequest{})
}

func (this Modules) Handle(request *api.HttpRequest, msbHandler api.MsbHandler) *api.HttpResponse {
	if request.RequestObject != nil {
		LOGGER.Info("request:", request.RequestObject)
	}

	m := make(map[string]interface{})
	var httpApis []string
	for _, v := range module.GetHttpApiHandlers() {
		httpApis = append(httpApis, v.Name())
	}
	m["http_api"] = httpApis

	return &api.HttpResponse{
		HttpStatus: http.StatusOK,
		HandleResult: api.HandleResult{
			Result: m,
		},
	}
}

func (this Modules) Name() string {
	return "builtin.modules"
}

func (this Modules) HttpMethods() []api.HttpMethod {
	return []api.HttpMethod{
		api.HttpMethodGet,
		api.HttpMethodPost,
	}
}

func (this Modules) UrlMapping() string {
	return "/builtin/modules"
}

var mgr api.HttpApiSimpleHandler = Modules{}

type ModuleJob struct {
}

func (m ModuleJob) CronConfig() string {
	return "*/1 * * * *"
}

func (m ModuleJob) Name() string {
	return "management.job"
}

func (m ModuleJob) Handler(msbHandler api.MsbHandler) (api.JobResultIndicator, error) {
	LOGGER.Info("job keep alive")

	return api.JobResultSuccess, nil
}

var mgrjob api.SchedulingJob = ModuleJob{}

func init() {
	module.RegisterHttpApiHandler(mgr)
	module.RegisterSchedulingJob(mgrjob)
}
