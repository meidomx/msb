package main

import (
	"net/http"
	"reflect"

	"github.com/meidomx/msb/api"
	"github.com/meidomx/msb/api/moduleapi"
	"github.com/meidomx/msb/kernel/process/database/postgres"
	"github.com/meidomx/msb/module"
)

type ProcessSample struct {
	api.DefaultHttpApiSimpleHandler
}

type ManagementRequest struct {
	Operation string
}

func (this ProcessSample) RequestType() reflect.Type {
	return reflect.TypeOf(ManagementRequest{})
}

func (this ProcessSample) Handle(request *api.HttpRequest, msbHandler api.MsbHandler) *api.HttpResponse {
	if request.RequestObject != nil {
		LOGGER_MODULE.Info("request:", request.RequestObject)
	}

	r, _ := msbHandler.CallProcess("tempprocess", &moduleapi.DatabaseQueryRequest{
		Operation:  "select",
		SQL:        "select * from princess_user",
		Parameters: [][]interface{}{},
	})
	LOGGER_MODULE.Info("call process result:", r)

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

func init() {

	if err := module.InstantiateBinding("builtin.database.provider.postgres", map[string]interface{}{
		"name":           "ds",
		"connect_string": "postgres://admin:admin@localhost:5432/princess",
	}, nil); err != nil {
		panic(err)
	}
	bd, err := module.GetBinding("ds").Bind(nil)
	if err != nil {
		panic(err)
	}
	if err := module.InstantiateService("builtin.database.operator.factory.postgres", map[string]interface{}{
		"name": "pg",
	}, bd); err != nil {
		panic(err)
	}

	module.RegisterProcess(TempProcess{
		pg: module.GetService("pg").(*postgres.PostgresDatabaseOperator),
	})

	ps := new(ProcessSample)
	ps.HandlerName = "sample.processsample"
	ps.RequestPayloadFormat = api.RequestFormatJson
	ps.ResponsePayloadFormat = api.ResponseFormatJson
	ps.HttpMethodList = []api.HttpMethod{api.HttpMethodGet, api.HttpMethodPost}
	ps.URL = "/api/sample/processsample"
	module.RegisterHttpApiHandler(ps)
}

type TempProcess struct {
	pg *postgres.PostgresDatabaseOperator
}

func (t TempProcess) Name() string {
	return "tempprocess"
}

func (t TempProcess) Call(param interface{}) (interface{}, error) {
	return t.pg.Handle(param)
}

var _ api.Process = TempProcess{}
