package main

import (
	"net/http"
	"reflect"

	"github.com/meidomx/msb/api"
	"github.com/meidomx/msb/module"
)

type HttpSampleHandler struct {
	api.DefaultHttpApiSimpleHandler
}

type Payload struct {
	Name string `json:"name"`
}

func (h *HttpSampleHandler) RequestType() reflect.Type {
	return reflect.TypeOf(Payload{})
}

func (h *HttpSampleHandler) Handle(request *api.HttpRequest, msbHandler api.MsbHandler) *api.HttpResponse {
	return &api.HttpResponse{
		HttpStatus: http.StatusOK,
		HandleResult: api.HandleResult{Result: map[string]interface{}{
			"name":    "hello",
			"age":     12,
			"already": true,
		}},
	}
}

var _ api.HttpApiSimpleHandler = new(HttpSampleHandler)

func init() {
	handler := new(HttpSampleHandler)
	handler.HandlerName = "sample.httpsample"
	handler.HttpMethodList = []api.HttpMethod{api.HttpMethodGet, api.HttpMethodPost}
	handler.RequestPayloadFormat = api.RequestFormatJson
	handler.ResponsePayloadFormat = api.ResponseFormatJson
	handler.URL = "/api/sample/httpsample"

	module.RegisterHttpApiHandler(handler)
}
